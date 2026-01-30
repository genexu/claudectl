package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"go.uber.org/fx"

	"claudectl/internal/domain"
	"claudectl/internal/loaders"
	"claudectl/internal/utils"
)

func RunNonInteractive(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	cfg Config,
	mcpLoader loaders.MCPLoader,
	commandLoader loaders.Loader[domain.Command],
	skillLoader loaders.Loader[domain.Skill],
	agentLoader loaders.Loader[domain.Agent],
	pluginLoader loaders.Loader[domain.Plugin],
	logger *utils.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Debug("running non-interactive mode")

			var capabilities []interface{}

			var scopesToLoad []domain.CapabilityScope
			switch cfg.ScopeFilter {
			case "user":
				scopesToLoad = []domain.CapabilityScope{domain.ScopeUser}
			case "project":
				scopesToLoad = []domain.CapabilityScope{domain.ScopeProject}
			case "all":
				scopesToLoad = []domain.CapabilityScope{domain.ScopeUser, domain.ScopeProject}
			default:
				scopesToLoad = []domain.CapabilityScope{domain.ScopeUser, domain.ScopeProject}
			}

			type loaderConfig struct {
				enabled bool
				load    func(domain.CapabilityScope)
			}

			loaders := []loaderConfig{
				{cfg.ListMCPs, func(scope domain.CapabilityScope) {
					loadCapabilities(mcpLoader, scope, &capabilities, logger, "MCP servers")
				}},
				{cfg.ListCommands, func(scope domain.CapabilityScope) {
					loadCapabilities(commandLoader, scope, &capabilities, logger, "commands")
				}},
				{cfg.ListSkills, func(scope domain.CapabilityScope) {
					loadCapabilities(skillLoader, scope, &capabilities, logger, "skills")
				}},
				{cfg.ListAgents, func(scope domain.CapabilityScope) {
					loadCapabilities(agentLoader, scope, &capabilities, logger, "agents")
				}},
				{cfg.ListPlugins, func(scope domain.CapabilityScope) {
					loadCapabilities(pluginLoader, scope, &capabilities, logger, "plugins")
				}},
			}

			for _, scope := range scopesToLoad {
				for _, loader := range loaders {
					if loader.enabled {
						loader.load(scope)
					}
				}
			}

			if cfg.JSONOutput {
				printJSON(capabilities)
			} else {
				printTable(capabilities)
			}

			return shutdowner.Shutdown()
		},
	})
}

type capabilityInfo struct {
	Name        string
	Scope       string
	Type        string
	Description string
}

func getCapabilityInfo(cap interface{}) capabilityInfo {
	switch v := cap.(type) {
	case *domain.MCPServer:
		return capabilityInfo{v.Name, string(v.Scope), string(v.Type), v.Description}
	case *domain.Command:
		return capabilityInfo{v.Name, string(v.Scope), string(v.Type), v.Description}
	case *domain.Skill:
		return capabilityInfo{v.Name, string(v.Scope), string(v.Type), v.Description}
	case *domain.Agent:
		return capabilityInfo{v.Name, string(v.Scope), "agent", v.Description}
	case *domain.Plugin:
		return capabilityInfo{v.Name, string(v.Scope), "plugin", v.Description}
	default:
		return capabilityInfo{}
	}
}

func printJSON(capabilities []interface{}) {
	data, err := json.MarshalIndent(capabilities, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func printTable(capabilities []interface{}) {
	if len(capabilities) == 0 {
		fmt.Println("No capabilities found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAME\tSCOPE\tTYPE\tDESCRIPTION")

	for _, cap := range capabilities {
		info := getCapabilityInfo(cap)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", info.Name, info.Scope, info.Type, info.Description)
	}

	w.Flush()
}
