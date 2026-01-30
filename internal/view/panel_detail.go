package view

import (
	"fmt"
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
		dp.SetContent("")
		return
	}

	var b strings.Builder

	b.WriteString(detailNameStyle.Render(vm.GetName()))
	b.WriteString(" " + detailScopeStyle.Render(fmt.Sprintf("[%s]", vm.GetScope())))

	if filepath := vm.GetFilePath(); filepath != "" {
		b.WriteString("\n\n" + detailFilepathStyle.Render(filepath))
	}

	if description := vm.GetDescription(); description != "" {
		b.WriteString("\n\n" + detailDescriptionStyle.Render(description))
	}

	if details := vm.RenderDetails(); len(details) > 0 {
		b.WriteString("\n")
		for _, detail := range details {
			b.WriteString("\n" + detail)
		}
	}

	if content := vm.GetContent(); content != "" {
		b.WriteString("\n\n" + content)
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
