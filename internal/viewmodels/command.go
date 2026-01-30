package viewmodels

import (
	"fmt"

	"claudectl/internal/domain"
)

type CommandViewModel struct {
	name        string
	description string
	scope       domain.CapabilityScope
	capType     domain.CapabilityType
	filePath    string
	content     string
}

func NewCommandViewModel(cmd *domain.Command) *CommandViewModel {
	return &CommandViewModel{
		name:        cmd.Name,
		description: cmd.Description,
		scope:       cmd.Scope,
		capType:     cmd.Type,
		filePath:    cmd.FilePath,
		content:     cmd.Content,
	}
}


func (vm *CommandViewModel) FilterValue() string {
	return vm.name
}

func (vm *CommandViewModel) Title() string {
	return vm.name
}

func (vm *CommandViewModel) Description() string {
	return fmt.Sprintf("[%s] %s", vm.scope, vm.description)
}


func (vm *CommandViewModel) RenderDetails() []string {
	// Metadata removed - return empty slice
	return []string{}
}


func (vm *CommandViewModel) GetName() string {
	return vm.name
}

func (vm *CommandViewModel) GetDescription() string {
	return vm.description
}

func (vm *CommandViewModel) GetScope() domain.CapabilityScope {
	return vm.scope
}

func (vm *CommandViewModel) GetType() domain.CapabilityType {
	return vm.capType
}

func (vm *CommandViewModel) GetFilePath() string {
	return vm.filePath
}

func (vm *CommandViewModel) GetContent() string {
	return vm.content
}
