package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListPanel struct {
	list  list.Model
	title string
}

func NewListPanel(items []list.Item, title string, width, height int) ListPanel {
	l := list.New(items, PanelListItemDelegate{}, width, height)
	l.Title = fmt.Sprintf("%s (%d)", title, len(items))
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return ListPanel{
		list:  l,
		title: title,
	}
}

func (lp ListPanel) Update(msg tea.Msg) (ListPanel, tea.Cmd) {
	var cmd tea.Cmd
	lp.list, cmd = lp.list.Update(msg)
	return lp, cmd
}

func (lp ListPanel) View() string {
	return lp.list.View()
}

func (lp *ListPanel) SetSize(width, height int) {
	lp.list.SetWidth(width)
	lp.list.SetHeight(height)
}

func (lp *ListPanel) SetItems(items []list.Item) {
	lp.list.SetItems(items)
	lp.list.Title = fmt.Sprintf("%s (%d)", lp.title, len(items))
}

func (lp ListPanel) SelectedItem() list.Item {
	return lp.list.SelectedItem()
}

func (lp ListPanel) ItemCount() int {
	return len(lp.list.Items())
}

func (lp *ListPanel) SelectFirst() {
	if lp.ItemCount() > 0 {
		lp.list.Select(0)
	}
}
