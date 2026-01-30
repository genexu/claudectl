package main

import (
	"flag"
	"fmt"
	"time"

	"go.uber.org/fx"

	"claudectl/internal/domain"
	"claudectl/internal/loaders"
	"claudectl/internal/utils"
	"claudectl/internal/view"
)

const version = "0.1.0"

type Config struct {
	Version      bool
	Debug        bool
	ListMCPs     bool
	ListCommands bool
	ListSkills   bool
	ListAgents   bool
	ListPlugins  bool
	ScopeFilter  string
	JSONOutput   bool
}

func (c Config) IsNonInteractive() bool {
	return c.ListMCPs || c.ListCommands || c.ListSkills || c.ListAgents || c.ListPlugins
}

func ParseFlags() Config {
	cfg := Config{}
	flag.BoolVar(&cfg.Version, "version", false, "Show version")
	flag.BoolVar(&cfg.Debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&cfg.ListMCPs, "list-mcps", false, "List MCP servers")
	flag.BoolVar(&cfg.ListCommands, "list-commands", false, "List commands")
	flag.BoolVar(&cfg.ListSkills, "list-skills", false, "List skills")
	flag.BoolVar(&cfg.ListAgents, "list-agents", false, "List agents")
	flag.BoolVar(&cfg.ListPlugins, "list-plugins", false, "List plugins")
	flag.StringVar(&cfg.ScopeFilter, "scope", "all", "Scope filter: user|project|all")
	flag.BoolVar(&cfg.JSONOutput, "json", false, "Output as JSON")
	flag.Parse()
	return cfg
}

func loadCapabilities[T any](
	loader loaders.Loader[T],
	scope domain.CapabilityScope,
	capabilities *[]interface{},
	logger *utils.Logger,
	capabilityType string,
) {
	caps, err := loader.Load(scope)
	if err != nil {
		logger.Warn(fmt.Sprintf("failed to load %s %s", scope, capabilityType), "error", err)
		return
	}
	for _, cap := range caps {
		c := cap
		*capabilities = append(*capabilities, &c)
	}
}

func main() {
	cfg := ParseFlags()

	if cfg.Version {
		fmt.Printf("claudectl version %s\n", version)
		return
	}

	options := []fx.Option{
		fx.Supply(cfg),
		fx.Provide(func(lc fx.Lifecycle, cfg Config) (*utils.Logger, error) {
			return utils.NewLogger(lc, cfg.Debug)
		}),
		fx.Provide(
			loaders.NewMCPLoader,
			loaders.NewCommandLoader,
			loaders.NewSkillLoader,
			loaders.NewAgentLoader,
			loaders.NewPluginLoader,
		),
		fx.Provide(view.NewModel),
		fx.StartTimeout(30 * time.Second),
		fx.StopTimeout(30 * time.Second),
		fx.NopLogger,
	}

	if cfg.IsNonInteractive() {
		options = append(options, fx.Invoke(RunNonInteractive))
	} else {
		options = append(options, fx.Invoke(RunTUI))
	}

	fx.New(options...).Run()
}
