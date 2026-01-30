package domain

// InstalledPluginsRegistry represents the installed_plugins.json file (domain model)
// e.g. ~/.claude/plugins/installed_plugins.json
type InstalledPluginsRegistry struct {
	Version int                              `json:"version"`
	Plugins map[string][]InstalledPluginInfo `json:"plugins"`
}

// InstalledPluginInfo represents a single plugin installation
type InstalledPluginInfo struct {
	Scope        CapabilityScope `json:"scope"`
	InstallPath  string          `json:"installPath"`
	Version      string          `json:"version"`
	InstalledAt  string          `json:"installedAt"`
	LastUpdated  string          `json:"lastUpdated"`
	GitCommitSha string          `json:"gitCommitSha"`
}
