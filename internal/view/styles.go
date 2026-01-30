package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

const (
	DefaultWidth  = 80
	DefaultHeight = 24

	LeftColumnWidthRatio = 0.40

	BorderSpacing     = 4
	TabBarHeight      = 6
	PanelSeparatorGap = 2
)

var (
	primaryColor   = lipgloss.Color("#6f03fc")
	secondaryColor = lipgloss.Color("#888888")
	bgColor        = lipgloss.Color("#1a1a1a")
	fgColor        = lipgloss.Color("#ffffff")
	dimColor       = lipgloss.Color("#555555")
	helpKeyColor   = lipgloss.Color("#666666")
	helpSepColor   = lipgloss.Color("#333333")

	activeTabStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 1)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Padding(0, 1)

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(primaryColor).
				Padding(1)

	unfocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(secondaryColor).
				Padding(1)

	panelTitleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Padding(1, 0)

	emptyStateStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Italic(true).
			Align(lipgloss.Center).
			Padding(2)

	detailNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Underline(true)

	detailScopeStyle = lipgloss.NewStyle().
				Italic(true).
				Foreground(secondaryColor)

	detailFilepathStyle = lipgloss.NewStyle().
				Foreground(fgColor).
				Italic(true)

	detailDescriptionStyle = lipgloss.NewStyle().
					Foreground(fgColor)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				PaddingLeft(2)

	unselectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(4)
)

func RenderTabs(activeTab TabType) string {
	tabs := []string{}

	for i := range TabCount {
		tab := TabType(i)
		style := inactiveTabStyle
		if tab == activeTab {
			style = activeTabStyle
		}
		tabLabel := style.Render("[" + string(rune('1'+i)) + "] " + tab.String())
		tabs = append(tabs, tabLabel)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func NewStyledHelp() help.Model {
	h := help.New()
	h.Styles.ShortKey = lipgloss.NewStyle().Foreground(helpKeyColor)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(secondaryColor)
	h.Styles.FullKey = lipgloss.NewStyle().Foreground(helpKeyColor)
	h.Styles.FullDesc = lipgloss.NewStyle().Foreground(secondaryColor)
	h.Styles.FullSeparator = lipgloss.NewStyle().Foreground(helpSepColor)
	return h
}
