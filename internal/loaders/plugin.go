package loaders

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"claudectl/internal/domain"
	"claudectl/internal/utils"
)

type PluginLoader struct {
	logger *slog.Logger
}

func NewPluginLoader(logger *utils.Logger) Loader[domain.Plugin] {
	logger.Debug("initializing plugin loader")
	return &PluginLoader{
		logger: logger.Logger,
	}
}

// PluginManifest represents the structure of .claude-plugin/plugin.json
type PluginManifest struct {
	Name        string              `json:"name"`
	Version     string              `json:"version"`
	Description string              `json:"description"`
	Author      domain.PluginAuthor `json:"author,omitempty"`
	Homepage    string              `json:"homepage,omitempty"`
	Repository  string              `json:"repository,omitempty"`
	License     string              `json:"license,omitempty"`
	Keywords    []string            `json:"keywords,omitempty"`
}

func (p *PluginLoader) Load(scope domain.CapabilityScope) ([]domain.Plugin, error) {
	// Get registry file path
	var registryPath string
	var err error

	if scope == domain.ScopeUser {
		registryPath, err = utils.GetUserInstalledPluginsFile()
	} else {
		registryPath, err = utils.GetProjectInstalledPluginsFile()
	}

	if err != nil {
		return nil, err
	}

	// Load registry (convert to domain registry)
	registry, err := p.loadInstalledPluginsRegistryDomain(registryPath)
	if err != nil {
		p.logger.Error("failed to load plugin registry", "path", registryPath, "error", err)
		return nil, err
	}

	if registry == nil {
		p.logger.Info("no plugin registry found", "scope", scope)
		return []domain.Plugin{}, nil
	}

	var capabilities []domain.Plugin

	// Iterate through registry entries
	for pluginKey, installations := range registry.Plugins {
		for _, installation := range installations {
			// Filter by scope
			if installation.Scope != scope {
				continue
			}

			// Load plugin from installPath
			plugin, err := p.loadPluginFromPathDomain(installation.InstallPath, installation.Version, scope, pluginKey)
			if err != nil {
				p.logger.Warn("failed to load plugin", "key", pluginKey, "path", installation.InstallPath, "error", err)
				continue
			}

			capabilities = append(capabilities, *plugin)
		}
	}

	p.logger.Info("discovered plugins from registry (domain)", "count", len(capabilities), "scope", scope)
	return capabilities, nil
}

// loadInstalledPluginsRegistryDomain reads and parses installed_plugins.json (domain model)
func (p *PluginLoader) loadInstalledPluginsRegistryDomain(registryPath string) (*domain.InstalledPluginsRegistry, error) {
	// Check if file exists
	if _, err := os.Stat(registryPath); os.IsNotExist(err) {
		return nil, nil // No registry = no plugins
	}

	data, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, err
	}

	var registry domain.InstalledPluginsRegistry
	if err := json.Unmarshal(data, &registry); err != nil {
		return nil, err
	}

	return &registry, nil
}

// loadPluginFromPathDomain loads a plugin from the given installation path (domain model)
func (p *PluginLoader) loadPluginFromPathDomain(installPath, version string, scope domain.CapabilityScope, pluginKey string) (*domain.Plugin, error) {
	manifestPath := filepath.Join(installPath, ".claude-plugin", "plugin.json")

	// Try to load manifest if it exists
	var manifest PluginManifest
	hasManifest := false

	if _, err := os.Stat(manifestPath); err == nil {
		// Manifest exists, load it
		data, err := os.ReadFile(manifestPath)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(data, &manifest); err != nil {
			return nil, err
		}
		hasManifest = true
	} else {
		// No manifest - extract name from plugin key
		atIndex := -1
		for i, c := range pluginKey {
			if c == '@' {
				atIndex = i
				break
			}
		}
		if atIndex > 0 {
			manifest.Name = pluginKey[:atIndex]
		} else {
			manifest.Name = pluginKey
		}
		manifest.Description = "Plugin installed from registry"
		p.logger.Debug("plugin manifest not found, using registry data", "key", pluginKey, "path", installPath)
	}

	// Create Plugin model (domain)
	plugin := &domain.Plugin{
		Capability: domain.Capability{
			Name:  manifest.Name,
			Description: manifest.Description,
			Type:  domain.TypePlugin,
			Scope: scope,
		},
		Version:     version,
		Author: domain.PluginAuthor{
			Name:  manifest.Author.Name,
			Email: manifest.Author.Email,
			URL:   manifest.Author.URL,
		},
		Homepage:   manifest.Homepage,
		Repository: manifest.Repository,
		License:    manifest.License,
		Keywords:   manifest.Keywords,
		Path:       installPath,
	}

	if hasManifest {
		p.logger.Debug("loaded plugin with manifest (domain)", "name", manifest.Name, "version", version, "scope", scope)
	} else {
		p.logger.Debug("loaded plugin without manifest (domain)", "name", manifest.Name, "version", version, "scope", scope)
	}

	return plugin, nil
}
