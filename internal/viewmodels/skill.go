package viewmodels

import (
	"fmt"

	"claudectl/internal/domain"
)

type SkillViewModel struct {
	// Common fields
	name        string
	description string
	scope       domain.CapabilityScope
	capType     domain.CapabilityType
	filePath    string
	content     string
}

func NewSkillViewModel(skill *domain.Skill) *SkillViewModel {
	return &SkillViewModel{
		name:        skill.Name,
		description: skill.Description,
		scope:       skill.Scope,
		capType:     skill.Type,
		filePath:    skill.FilePath,
		content:     skill.Content,
	}
}

func (vm *SkillViewModel) FilterValue() string {
	return vm.name
}

func (vm *SkillViewModel) Title() string {
	return vm.name
}

func (vm *SkillViewModel) Description() string {
	return fmt.Sprintf("[%s] %s", vm.scope, vm.description)
}

func (vm *SkillViewModel) RenderDetails() []string {
	// No type-specific details to display
	return []string{}
}

func (vm *SkillViewModel) GetName() string {
	return vm.name
}

func (vm *SkillViewModel) GetDescription() string {
	return vm.description
}

func (vm *SkillViewModel) GetScope() domain.CapabilityScope {
	return vm.scope
}

func (vm *SkillViewModel) GetType() domain.CapabilityType {
	return vm.capType
}

func (vm *SkillViewModel) GetFilePath() string {
	return vm.filePath
}

func (vm *SkillViewModel) GetContent() string {
	return vm.content
}
