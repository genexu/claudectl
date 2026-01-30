package view

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type PanelListItemDelegate struct{}

func (d PanelListItemDelegate) Height() int { return 1 }

func (d PanelListItemDelegate) Spacing() int { return 0 }

func (d PanelListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d PanelListItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	title := listItem.FilterValue()
	if index == m.Index() {
		icon := selectedItemIconStyle.Render(SymbolSelected + " ")
		content := selectedItemStyle.Render(icon + title)
		fmt.Fprint(w, content)
	} else {
		icon := unselectedItemIconStyle.Render(SymbolUnselected + " ")
		content := unselectedItemStyle.Render(icon + title)
		fmt.Fprint(w, content)
	}
}
