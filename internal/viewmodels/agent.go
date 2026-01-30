package viewmodels

import (
	"fmt"

	"claudectl/internal/domain"
)

type AgentViewModel struct {
	name        string
	description string
	scope       domain.CapabilityScope
	capType     domain.CapabilityType
	filePath    string
	content     string
}

func NewAgentViewModel(agent *domain.Agent) *AgentViewModel {
	return &AgentViewModel{
		name:        agent.Name,
		description: agent.Description,
		scope:       agent.Scope,
		capType:     agent.Type,
		filePath:    agent.FilePath,
		content:     agent.Content,
	}
}


func (vm *AgentViewModel) FilterValue() string {
	return vm.name
}

func (vm *AgentViewModel) Title() string {
	return vm.name
}

func (vm *AgentViewModel) Description() string {
	return fmt.Sprintf("[%s] %s", vm.scope, vm.description)
}


func (vm *AgentViewModel) RenderDetails() []string {
	// No type-specific details to display
	return []string{}
}


func (vm *AgentViewModel) GetName() string {
	return vm.name
}

func (vm *AgentViewModel) GetDescription() string {
	return vm.description
}

func (vm *AgentViewModel) GetScope() domain.CapabilityScope {
	return vm.scope
}

func (vm *AgentViewModel) GetType() domain.CapabilityType {
	return vm.capType
}

func (vm *AgentViewModel) GetFilePath() string {
	return vm.filePath
}

func (vm *AgentViewModel) GetContent() string {
	return vm.content
}
