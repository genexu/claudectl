package view

import (
	"strings"

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

// Minimal symbol set - Clean and functional
const (
	// Navigation
	SymbolSelected   = "●"
	SymbolUnselected = "○"
	SymbolArrow      = "▶"

	// Status
	SymbolCheck      = "✓"
	SymbolCross      = "✗"
	SymbolWarning    = "⚠"

	// Minimal decoration
	SymbolDot        = "•"
	SymbolPrompt     = "❯"
)

var (
	// Navy Blue color palette - Professional & Sophisticated
	primaryColor     = lipgloss.Color("#1E40AF") // Navy Blue
	primaryLight     = lipgloss.Color("#3B82F6") // Medium Blue
	primaryBright    = lipgloss.Color("#60A5FA") // Light Blue

	// Clean backgrounds - Dark mode OLED
	bgColor          = lipgloss.Color("#0a0e14") // Deep black
	bgSurface        = lipgloss.Color("#1a1f2e") // Elevated surface
	bgHighlight      = lipgloss.Color("#1e293b") // Highlight surface

	// High contrast text
	fgColor          = lipgloss.Color("#F8FAFC") // Off white
	textMuted        = lipgloss.Color("#94a3b8") // Slate muted
	textDim          = lipgloss.Color("#64748b") // Slate dim

	// Semantic colors (minimal)
	successColor     = lipgloss.Color("#10b981") // Green (for success only)
	warningColor     = lipgloss.Color("#f59e0b") // Amber
	errorColor       = lipgloss.Color("#ef4444") // Red
	infoColor        = lipgloss.Color("#3B82F6") // Blue (info)

	// UI element colors
	borderColor      = lipgloss.Color("#334155") // Slate border
	borderActive     = lipgloss.Color("#3B82F6") // Active blue border

	// Tab styles - Clean block design with Navy Blue
	activeTabStyle = lipgloss.NewStyle().
		Foreground(fgColor).
		Background(primaryColor).
		Bold(true).
		Padding(0, 2).
		MarginRight(1)

	inactiveTabStyle = lipgloss.NewStyle().
		Foreground(textDim).
		Background(bgSurface).
		Padding(0, 2).
		MarginRight(1)

	tabBarStyle = lipgloss.NewStyle().
		Background(bgColor).
		Padding(1, 0)

	// Clean border styles
	focusedBorderStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderActive).
		Padding(1)

	unfocusedBorderStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(1)

	// Panel title - Clean and minimal
	panelTitleStyle = lipgloss.NewStyle().
		Foreground(primaryLight).
		Bold(true).
		Padding(0, 1)

	// Help section
	helpStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		Background(bgSurface).
		Padding(0, 1)

	// Empty state
	emptyStateStyle = lipgloss.NewStyle().
		Foreground(textDim).
		Italic(true).
		Align(lipgloss.Center).
		Padding(4, 2)

	emptyStateIconStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Align(lipgloss.Center)

	// Detail panel styles - Clean typography with Navy Blue
	detailNameStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryBright)

	detailScopeStyle = lipgloss.NewStyle().
		Foreground(primaryLight).
		Bold(true)

	detailFilepathStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		Italic(false)

	detailDescriptionStyle = lipgloss.NewStyle().
		Foreground(fgColor).
		Padding(0, 0, 1, 0)

	detailSectionHeaderStyle = lipgloss.NewStyle().
		Foreground(primaryLight).
		Bold(true).
		Padding(1, 0, 0, 0)

	detailValueStyle = lipgloss.NewStyle().
		Foreground(fgColor)

	// List item styles - Minimal and clean
	selectedItemStyle = lipgloss.NewStyle().
		Foreground(fgColor).
		Background(bgHighlight).
		Bold(true).
		Padding(0, 1).
		MarginLeft(1).
		MarginRight(1)

	unselectedItemStyle = lipgloss.NewStyle().
		Foreground(textMuted).
		Padding(0, 1).
		MarginLeft(1).
		MarginRight(1)

	selectedItemIconStyle = lipgloss.NewStyle().
		Foreground(primaryBright).
		Bold(true)

	unselectedItemIconStyle = lipgloss.NewStyle().
		Foreground(borderColor)

	// Status badges - Clean and functional with Navy Blue
	userScopeBadgeStyle = lipgloss.NewStyle().
		Foreground(fgColor).
		Background(primaryLight).
		Bold(true).
		Padding(0, 1)

	projectScopeBadgeStyle = lipgloss.NewStyle().
		Foreground(fgColor).
		Background(primaryColor).
		Bold(true).
		Padding(0, 1)

	// Divider style - Simple
	dividerStyle = lipgloss.NewStyle().
		Foreground(borderColor)

	// Status indicators
	statusSuccessStyle = lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true)

	statusWarningStyle = lipgloss.NewStyle().
		Foreground(warningColor).
		Bold(true)

	statusErrorStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Bold(true)

	statusInfoStyle = lipgloss.NewStyle().
		Foreground(infoColor).
		Bold(true)

	// Terminal prompt style - Navy Blue
	terminalPromptStyle = lipgloss.NewStyle().
		Foreground(primaryBright).
		Bold(true)
)

// RenderTabs renders the tab bar with clean block design
func RenderTabs(activeTab TabType) string {
	tabs := []string{}

	for i := 0; i < TabCount; i++ {
		tab := TabType(i)
		style := inactiveTabStyle
		if tab == activeTab {
			style = activeTabStyle
		}
		tabLabel := style.Render("[" + string(rune('1'+i)) + "] " + tab.String())
		tabs = append(tabs, tabLabel)
	}

	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	return tabBarStyle.Render(tabBar)
}

// NewStyledHelp creates a help model with clean styling
func NewStyledHelp() help.Model {
	h := help.New()
	h.Styles.ShortKey = lipgloss.NewStyle().Foreground(textDim)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(textMuted)
	h.Styles.FullKey = lipgloss.NewStyle().Foreground(textDim)
	h.Styles.FullDesc = lipgloss.NewStyle().Foreground(textMuted)
	h.Styles.FullSeparator = lipgloss.NewStyle().Foreground(borderColor)
	return h
}

// RenderScopeBadge renders a clean scope badge
func RenderScopeBadge(scope string) string {
	switch scope {
	case "user", "User":
		return userScopeBadgeStyle.Render(" USER ")
	case "project", "Project":
		return projectScopeBadgeStyle.Render(" PROJECT ")
	default:
		return statusInfoStyle.Render(" " + scope + " ")
	}
}

// RenderStatusIndicator renders a status with minimal symbols
func RenderStatusIndicator(status string) string {
	switch status {
	case "success", "enabled", "active":
		return statusSuccessStyle.Render(SymbolCheck + " " + status)
	case "warning", "pending":
		return statusWarningStyle.Render(SymbolWarning + " " + status)
	case "error", "disabled", "failed":
		return statusErrorStyle.Render(SymbolCross + " " + status)
	default:
		return statusInfoStyle.Render(SymbolDot + " " + status)
	}
}

// RenderDivider renders a simple horizontal divider
func RenderDivider(width int) string {
	return dividerStyle.Render(strings.Repeat("─", width))
}

// RenderTerminalPrompt renders a clean terminal-style prompt
func RenderTerminalPrompt(text string) string {
	prompt := terminalPromptStyle.Render(SymbolPrompt + " ")
	return prompt + text
}
