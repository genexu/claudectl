package view

import "claudectl/internal/domain"

type TabType int

const (
	MCPsTab TabType = iota
	CommandsTab
	SkillsTab
	PluginsTab
	AgentsTab
)

func (t TabType) String() string {
	switch t {
	case MCPsTab:
		return "MCPs"
	case CommandsTab:
		return "Commands"
	case SkillsTab:
		return "Skills"
	case PluginsTab:
		return "Plugins"
	case AgentsTab:
		return "Agents"
	default:
		return "Unknown"
	}
}

const TabCount = 5

func (t TabType) NextTab() TabType {
	return TabType((int(t) + 1) % TabCount)
}

func (t TabType) PrevTab() TabType {
	return TabType((int(t) - 1 + TabCount) % TabCount)
}

func (t TabType) ToCapabilityType() domain.CapabilityType {
	switch t {
	case MCPsTab:
		return domain.TypeMCP
	case CommandsTab:
		return domain.TypeCommand
	case SkillsTab:
		return domain.TypeSkill
	case AgentsTab:
		return domain.TypeAgent
	case PluginsTab:
		return domain.TypePlugin
	default:
		return domain.TypeMCP
	}
}
