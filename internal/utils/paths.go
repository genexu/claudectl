package utils

import (
	"claudectl/internal/domain"
	"os"
	"path/filepath"
)

func GetScopeBaseDir(scope domain.CapabilityScope) (string, error) {
	if scope == domain.ScopeUser {
		return GetUserClaudeDir()
	}

	return GetProjectClaudeDir()
}

// i.e., ~/.claude
func GetUserClaudeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude"), nil
}

// i.e., /path/to/project/.claude
func GetProjectClaudeDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, ".claude"), nil
}

// i.e., ~/.claude/plugins
func GetUserPluginsDir() (string, error) {
	userDir, err := GetUserClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userDir, "plugins"), nil
}

// i.e., /path/to/project/.claude/plugins
func GetProjectPluginsDir() (string, error) {
	projectDir, err := GetProjectClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(projectDir, "plugins"), nil
}

// e.g., /home/user/.claude.json
func GetUserConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude.json"), nil
}

// e.g., /path/to/project/.claude/.mcp.json
func GetProjectMCPConfigFile() (string, error) {
	projectDir, err := GetProjectClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(projectDir, ".mcp.json"), nil
}

// e.g., /home/user/.claude/plugins/installed_plugins.json
func GetUserInstalledPluginsFile() (string, error) {
	pluginsDir, err := GetUserPluginsDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(pluginsDir, "installed_plugins.json"), nil
}

// e.g., /path/to/project/.claude/plugins/installed_plugins.json
func GetProjectInstalledPluginsFile() (string, error) {
	pluginsDir, err := GetProjectPluginsDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(pluginsDir, "installed_plugins.json"), nil
}
