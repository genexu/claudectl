package viewmodels

import (
	"fmt"

	"claudectl/internal/domain"
)

type PluginViewModel struct {
	// Common fields
	name        string
	description string
	scope       domain.CapabilityScope
	capType     domain.CapabilityType
	path        string

	// Plugin-specific fields
	version    string
	authorName string
	license    string
}

func NewPluginViewModel(plugin *domain.Plugin) *PluginViewModel {
	return &PluginViewModel{
		name:        plugin.Name,
		description: plugin.Description,
		scope:       plugin.Scope,
		capType:     plugin.Type,
		path:        plugin.Path,
		version:     plugin.Version,
		authorName:  plugin.Author.Name,
		license:     plugin.License,
	}
}

func (vm *PluginViewModel) FilterValue() string {
	return vm.name
}

func (vm *PluginViewModel) Title() string {
	return vm.name
}

func (vm *PluginViewModel) Description() string {
	return fmt.Sprintf("[%s] %s", vm.scope, vm.description)
}

func (vm *PluginViewModel) RenderDetails() []string {
	details := []string{
		fmt.Sprintf("Version: %s", vm.version),
	}

	if vm.authorName != "" {
		details = append(details, fmt.Sprintf("Author: %s", vm.authorName))
	}

	if vm.license != "" {
		details = append(details, fmt.Sprintf("License: %s", vm.license))
	}

	return details
}

func (vm *PluginViewModel) GetName() string {
	return vm.name
}

func (vm *PluginViewModel) GetDescription() string {
	return vm.description
}

func (vm *PluginViewModel) GetScope() domain.CapabilityScope {
	return vm.scope
}

func (vm *PluginViewModel) GetType() domain.CapabilityType {
	return vm.capType
}

func (vm *PluginViewModel) GetFilePath() string {
	return vm.path
}

// GetContent returns the markdown content (plugins don't have content, return empty)
func (vm *PluginViewModel) GetContent() string {
	return ""
}
