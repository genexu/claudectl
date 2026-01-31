package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"claudectl/internal/domain"
	"claudectl/internal/loaders"
	"claudectl/internal/utils"
	"claudectl/internal/viewmodels"
)

type PanelType int

const (
	UserPanel PanelType = iota
	ProjectPanel
	DetailPanelFocus
)

type panelDimensions struct {
	leftColumnWidth    int
	detailWidth        int
	panelHeight        int
	userPanelHeight    int
	projectPanelHeight int
}

type Model struct {
	logger        *utils.Logger
	mcpLoader     loaders.MCPLoader
	commandLoader loaders.Loader[domain.Command]
	skillLoader   loaders.Loader[domain.Skill]
	agentLoader   loaders.Loader[domain.Agent]
	pluginLoader  loaders.Loader[domain.Plugin]

	activeTab   TabType
	activePanel PanelType
	width       int
	height      int

	help help.Model
	keys KeyMap

	userListPanel    ListPanel
	projectListPanel ListPanel
	detailPanel      DetailPanel

	userCapabilities    []viewmodels.CapabilityViewModel
	projectCapabilities []viewmodels.CapabilityViewModel

	program *tea.Program
}

func NewModel(
	logger *utils.Logger,
	mcpLoader loaders.MCPLoader,
	commandLoader loaders.Loader[domain.Command],
	skillLoader loaders.Loader[domain.Skill],
	agentLoader loaders.Loader[domain.Agent],
	pluginLoader loaders.Loader[domain.Plugin],
) *Model {
	model := &Model{
		logger:        logger,
		mcpLoader:     mcpLoader,
		commandLoader: commandLoader,
		skillLoader:   skillLoader,
		agentLoader:   agentLoader,
		pluginLoader:  pluginLoader,
		activeTab:     MCPsTab,
		activePanel:   UserPanel,
		width:  DefaultWidth,
		height: DefaultHeight,
		keys:   DefaultKeyMap(),
		help:   NewStyledHelp(),
	}

	model.loadCapabilities()

	dims := model.calculatePanelDimensions()

	capType := model.activeTab.ToCapabilityType()
	userItems := model.filterByType(model.userCapabilities, capType)
	projectItems := model.filterByType(model.projectCapabilities, capType)

	model.userListPanel = NewListPanel(userItems, "User",
		dims.leftColumnWidth-4, dims.userPanelHeight-4)
	model.projectListPanel = NewListPanel(projectItems, "Project",
		dims.leftColumnWidth-4, dims.projectPanelHeight-4)
	model.detailPanel = NewDetailPanel(dims.detailWidth-4, dims.panelHeight-4)

	model.updateDetailPanel()

	return model
}

func (m *Model) SetProgram(p *tea.Program) {
	m.program = p
}

func (m *Model) calculatePanelDimensions() panelDimensions {
	leftColumnWidth := int(float64(m.width) * LeftColumnWidthRatio)
	detailWidth := m.width - leftColumnWidth - BorderSpacing
	panelHeight := m.height - TabBarHeight
	userPanelHeight := (panelHeight - PanelSeparatorGap) / 2
	projectPanelHeight := panelHeight - userPanelHeight - PanelSeparatorGap

	return panelDimensions{
		leftColumnWidth:    leftColumnWidth,
		detailWidth:        detailWidth,
		panelHeight:        panelHeight,
		userPanelHeight:    userPanelHeight,
		projectPanelHeight: projectPanelHeight,
	}
}

func (m *Model) loadCapabilities() {
	loadFromLoader(m.mcpLoader, &m.userCapabilities, &m.projectCapabilities, m.logger)
	loadFromLoader(m.commandLoader, &m.userCapabilities, &m.projectCapabilities, m.logger)
	loadFromLoader(m.skillLoader, &m.userCapabilities, &m.projectCapabilities, m.logger)
	loadFromLoader(m.agentLoader, &m.userCapabilities, &m.projectCapabilities, m.logger)
	loadFromLoader(m.pluginLoader, &m.userCapabilities, &m.projectCapabilities, m.logger)

	if m.logger != nil {
		m.logger.Info("loaded capabilities",
			"user_capabilities", len(m.userCapabilities),
			"project_capabilities", len(m.projectCapabilities))
	}
}

func loadFromLoader[T any](
	loader loaders.Loader[T],
	userCaps *[]viewmodels.CapabilityViewModel,
	projectCaps *[]viewmodels.CapabilityViewModel,
	logger *utils.Logger,
) {
	loadScope := func(scope domain.CapabilityScope, caps *[]viewmodels.CapabilityViewModel, scopeName string) {
		if items, err := loader.Load(scope); err == nil {
			for _, cap := range items {
				vm, err := viewmodels.ToDomainViewModel(cap)
				if err != nil {
					if logger != nil {
						logger.Warn("failed to convert capability to view model", "error", err)
					}
					continue
				}
				*caps = append(*caps, vm)
			}
		} else if logger != nil {
			logger.Warn("failed to load "+scopeName+" capabilities", "error", err)
		}
	}

	loadScope(domain.ScopeUser, userCaps, "user")
	loadScope(domain.ScopeProject, projectCaps, "project")
}

func (m *Model) updateListsForCurrentTab() {
	capType := m.activeTab.ToCapabilityType()
	userItems := m.filterByType(m.userCapabilities, capType)
	projectItems := m.filterByType(m.projectCapabilities, capType)

	m.userListPanel.SetItems(userItems)
	m.projectListPanel.SetItems(projectItems)
}

func (m *Model) filterByType(items []viewmodels.CapabilityViewModel, capType domain.CapabilityType) []list.Item {
	var filtered []list.Item
	for _, item := range items {
		if item.GetType() == capType {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func (m *Model) updateDetailPanel() {
	selectedItem := m.userListPanel.SelectedItem()
	if m.activePanel != UserPanel {
		selectedItem = m.projectListPanel.SelectedItem()
	}

	if selectedItem == nil {
		m.detailPanel.SetContent("")
		return
	}

	m.detailPanel.Render(selectedItem)
}

func (m *Model) selectFirstInActivePanel() {
	if m.activePanel == UserPanel {
		m.userListPanel.SelectFirst()
	} else {
		m.projectListPanel.SelectFirst()
	}
	m.updateDetailPanel()
}

func (m *Model) switchToTab(tab TabType) tea.Cmd {
	m.activeTab = tab
	m.updateListsForCurrentTab()
	m.selectFirstInActivePanel()
	if m.logger != nil {
		m.logger.Debug("switched tab", "tab", tab.String())
	}
	return nil
}

func (m Model) ActiveTab() TabType {
	return m.activeTab
}

func (m Model) ActivePanel() PanelType {
	return m.activePanel
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, m.keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
		}

		tabKeys := []struct {
			binding key.Binding
			tab     TabType
		}{
			{m.keys.Tab1, MCPsTab},
			{m.keys.Tab2, CommandsTab},
			{m.keys.Tab3, SkillsTab},
			{m.keys.Tab4, PluginsTab},
			{m.keys.Tab5, AgentsTab},
		}
		for _, tk := range tabKeys {
			if key.Matches(msg, tk.binding) {
				return m, m.switchToTab(tk.tab)
			}
		}

		if key.Matches(msg, m.keys.SwitchPanel) {
			switch m.activePanel {
			case UserPanel:
				m.activePanel = ProjectPanel
				m.projectListPanel.SelectFirst()
			case ProjectPanel:
				m.activePanel = DetailPanelFocus
			case DetailPanelFocus:
				m.activePanel = UserPanel
				m.userListPanel.SelectFirst()
			}
			m.updateDetailPanel()
			if m.logger != nil {
				m.logger.Debug("switched panel focus", "panel", m.activePanel)
			}
			return m, nil
		}

		handleTabCycle := func(nextTab TabType, direction string) (tea.Model, tea.Cmd) {
			m.activeTab = nextTab
			m.updateListsForCurrentTab()
			m.selectFirstInActivePanel()
			if m.logger != nil {
				m.logger.Debug("cycled tab "+direction, "tab", m.activeTab.String())
			}
			return m, nil
		}

		if key.Matches(msg, m.keys.NextTab) {
			return handleTabCycle(m.activeTab.NextTab(), "forward")
		}
		if key.Matches(msg, m.keys.PrevTab) {
			return handleTabCycle(m.activeTab.PrevTab(), "backward")
		}

		if key.Matches(msg, m.keys.Up) || key.Matches(msg, m.keys.Down) ||
			key.Matches(msg, m.keys.PageUp) || key.Matches(msg, m.keys.PageDown) {
			var cmd tea.Cmd
			switch m.activePanel {
			case UserPanel:
				m.userListPanel, cmd = m.userListPanel.Update(msg)
				m.updateDetailPanel()
			case ProjectPanel:
				m.projectListPanel, cmd = m.projectListPanel.Update(msg)
				m.updateDetailPanel()
			case DetailPanelFocus:
				cmd = m.detailPanel.Update(msg)
			}
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		dims := m.calculatePanelDimensions()

		m.userListPanel.SetSize(dims.leftColumnWidth-4, dims.userPanelHeight-4)
		m.projectListPanel.SetSize(dims.leftColumnWidth-4, dims.projectPanelHeight-4)
		m.detailPanel.SetSize(dims.detailWidth-4, dims.panelHeight-4)

		if m.logger != nil {
			m.logger.Debug("window resized", "width", m.width, "height", m.height)
		}
	}

	return m, nil
}

func (m Model) View() string {
	dims := m.calculatePanelDimensions()
	tabBar := RenderTabs(m.activeTab)

	userBorderStyle := unfocusedBorderStyle
	projectBorderStyle := unfocusedBorderStyle
	detailBorderStyle := unfocusedBorderStyle

	switch m.activePanel {
	case UserPanel:
		userBorderStyle = focusedBorderStyle
	case ProjectPanel:
		projectBorderStyle = focusedBorderStyle
	case DetailPanelFocus:
		detailBorderStyle = focusedBorderStyle
	}

	userPanel := userBorderStyle.
		Width(dims.leftColumnWidth).
		Height(dims.userPanelHeight).
		Render(m.userListPanel.View())

	projectPanel := projectBorderStyle.
		Width(dims.leftColumnWidth).
		Height(dims.projectPanelHeight).
		Render(m.projectListPanel.View())

	leftColumn := lipgloss.JoinVertical(lipgloss.Left, userPanel, projectPanel)

	detailPanel := detailBorderStyle.
		Width(dims.detailWidth).
		Height(dims.panelHeight).
		Render(m.detailPanel.View())

	panels := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, detailPanel)

	helpView := m.help.View(m.keys)
	help := helpStyle.Width(m.width).Render(helpView)

	return lipgloss.JoinVertical(lipgloss.Left, tabBar, panels, help)
}
