package loaders

import (
	"encoding/json"
	"log/slog"
	"os"

	"claudectl/internal/domain"
	"claudectl/internal/utils"
)

type MCPLoader interface {
	Loader[domain.MCPServer] // Embed unified loader interface
}

type mcpLoaderImpl struct {
	logger *slog.Logger
}

func NewMCPLoader(logger *utils.Logger) MCPLoader {
	logger.Debug("initializing MCP loader")
	return &mcpLoaderImpl{logger: logger.Logger}
}

// MCPServerConfig represents the structure of MCP server configuration
type MCPServerConfig struct {
	MCPType string            `json:"type"`
	Command string            `json:"command,omitempty"`
	Args    []string          `json:"args,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
	Url     string            `json:"url,omitempty"`
}

// ClaudeConfig represents the structure of .claude.json or .mcp.json
type ClaudeConfig struct {
	MCPServers map[string]MCPServerConfig `json:"mcpServers"`
}

func (m *mcpLoaderImpl) Load(scope domain.CapabilityScope) ([]domain.MCPServer, error) {
	// Get config path based on scope
	var pathGetter func() (string, error)

	if scope == domain.ScopeUser {
		pathGetter = utils.GetUserConfigFile
	} else {
		pathGetter = utils.GetProjectMCPConfigFile
	}

	configPath, err := pathGetter()
	if err != nil {
		m.logger.Error("failed to get config path", "scope", scope, "error", err)
		return nil, err
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		m.logger.Info("config file not found, skipping", "path", configPath)
		return []domain.MCPServer{}, nil
	}

	// Read file
	data, err := os.ReadFile(configPath)
	if err != nil {
		m.logger.Error("failed to read config file", "path", configPath, "error", err)
		return nil, err
	}

	// Parse JSON
	var config ClaudeConfig
	if err := json.Unmarshal(data, &config); err != nil {
		m.logger.Error("failed to parse config JSON", "path", configPath, "error", err)
		return nil, err
	}

	// Convert to domain MCPServer models
	var capabilities []domain.MCPServer
	for name, serverConfig := range config.MCPServers {
		server := *domain.NewMCPServer(domain.MCPServerParams{
			Name:    name,
			Scope:   scope,
			Command: serverConfig.Command,
			Args:    serverConfig.Args,
			Env:     serverConfig.Env,
			MCPType: serverConfig.MCPType,
			Url:     serverConfig.Url,
		})
		capabilities = append(capabilities, server)
		m.logger.Debug("loaded MCP server", "name", name, "scope", scope)
	}

	m.logger.Info("loaded MCP servers (domain)", "count", len(capabilities), "scope", scope, "path", configPath)
	return capabilities, nil
}
