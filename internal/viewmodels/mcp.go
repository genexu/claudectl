package viewmodels

import (
	"fmt"

	"claudectl/internal/domain"
)

type MCPServerViewModel struct {
	// Common fields
	name        string
	description string
	scope       domain.CapabilityScope
	capType     domain.CapabilityType

	// MCP-specific fields
	command string
	args    []string
	env     map[string]string
	mcpType string
	url     string
}

func NewMCPServerViewModel(server *domain.MCPServer) *MCPServerViewModel {
	return &MCPServerViewModel{
		name:        server.Name,
		description: server.Description,
		scope:       server.Scope,
		capType:     server.Type,
		command:     server.Command,
		args:        server.Args,
		env:         server.Env,
		mcpType:     server.MCPType,
		url:         server.Url,
	}
}

func (vm *MCPServerViewModel) FilterValue() string {
	return vm.name
}

func (vm *MCPServerViewModel) Title() string {
	return vm.name
}

func (vm *MCPServerViewModel) Description() string {
	return fmt.Sprintf("[%s] %s", vm.scope, vm.description)
}

func (vm *MCPServerViewModel) RenderDetails() []string {
	details := []string{}

	if vm.mcpType != "" {
		details = append(details, fmt.Sprintf("Type: %s", vm.mcpType))
	}

	if vm.url != "" {
		details = append(details, fmt.Sprintf("URL: %s", vm.url))
	}

	if vm.command != "" {
		details = append(details, fmt.Sprintf("Command: %s", vm.command))
	}

	if len(vm.args) > 0 {
		details = append(details, "Arguments:")
		for _, arg := range vm.args {
			details = append(details, fmt.Sprintf(" %s", arg))
		}
	}

	if len(vm.env) > 0 {
		details = append(details, "Environment:\n")
		for k, v := range vm.env {
			details = append(details, fmt.Sprintf("  %s=%s\n", k, v))
		}
	}

	return details
}

func (vm *MCPServerViewModel) GetName() string {
	return vm.name
}

func (vm *MCPServerViewModel) GetDescription() string {
	return vm.description
}

func (vm *MCPServerViewModel) GetScope() domain.CapabilityScope {
	return vm.scope
}

func (vm *MCPServerViewModel) GetType() domain.CapabilityType {
	return vm.capType
}

// GetFilePath returns the file path (MCP servers don't have a file path, return empty)
func (vm *MCPServerViewModel) GetFilePath() string {
	return ""
}

// GetContent returns the markdown content (MCP servers don't have content, return empty)
func (vm *MCPServerViewModel) GetContent() string {
	return ""
}
