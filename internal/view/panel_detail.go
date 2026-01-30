package view

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"

	"claudectl/internal/viewmodels"
)

type DetailPanel struct {
	viewport viewport.Model
	ready    bool
}

func NewDetailPanel(width, height int) DetailPanel {
	vp := viewport.New(width, height)
	vp.Style = lipgloss.NewStyle()

	return DetailPanel{
		viewport: vp,
		ready:    true,
	}
}

func (dp *DetailPanel) SetSize(width, height int) {
	dp.viewport.Width = width
	dp.viewport.Height = height
}

func (dp *DetailPanel) SetContent(content string) {
	if dp.viewport.Width > 0 {
		content = wordwrap.String(content, dp.viewport.Width)
	}
	dp.viewport.SetContent(content)
}

func (dp *DetailPanel) Render(item interface{}) {
	vm, ok := item.(viewmodels.CapabilityViewModel)
	if !ok {
		emptyIcon := emptyStateIconStyle.Render(SymbolDot)
		emptyText := emptyStateStyle.Render("Select a capability to view details")
		dp.SetContent(emptyIcon + "\n\n" + emptyText)
		return
	}

	var b strings.Builder

	// Header: Name and Scope Badge (clean inline)
	nameAndBadge := detailNameStyle.Render(vm.GetName()) + " " + RenderScopeBadge(string(vm.GetScope()))
	b.WriteString(nameAndBadge)
	b.WriteString("\n")

	// Filepath section with terminal prompt
	if filepath := vm.GetFilePath(); filepath != "" {
		b.WriteString("\n")
		b.WriteString(detailSectionHeaderStyle.Render("Location"))
		b.WriteString("\n")
		b.WriteString(RenderTerminalPrompt(detailFilepathStyle.Render(filepath)))
		b.WriteString("\n")
	}

	// Description section
	if description := vm.GetDescription(); description != "" {
		b.WriteString("\n")
		b.WriteString(detailSectionHeaderStyle.Render("Description"))
		b.WriteString("\n")
		b.WriteString(detailDescriptionStyle.Render(description))
		b.WriteString("\n")
	}

	// Additional details section
	if details := vm.RenderDetails(); len(details) > 0 {
		b.WriteString("\n")
		b.WriteString(detailSectionHeaderStyle.Render("Details"))
		b.WriteString("\n")
		for _, detail := range details {
			b.WriteString(detailValueStyle.Render(detail))
			b.WriteString("\n")
		}
	}

	// Content section with clean divider
	if content := vm.GetContent(); content != "" {
		b.WriteString("\n")
		b.WriteString(RenderDivider(dp.viewport.Width))
		b.WriteString("\n\n")
		b.WriteString(content)
	}

	dp.SetContent(b.String())
}

func (dp *DetailPanel) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	dp.viewport, cmd = dp.viewport.Update(msg)
	return cmd
}

func (dp DetailPanel) View() string {
	if !dp.ready {
		return emptyStateStyle.Width(dp.viewport.Width).Height(dp.viewport.Height).Render("Select a capability")
	}
	return dp.viewport.View()
}
