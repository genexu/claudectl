package view

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	PageUp   key.Binding
	PageDown key.Binding

	SwitchListPanel key.Binding

	Tab1 key.Binding
	Tab2 key.Binding
	Tab3 key.Binding
	Tab4 key.Binding
	Tab5 key.Binding

	Help key.Binding
	Quit key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
		SwitchListPanel: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch panels"),
		),
		Tab1: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("", ""),
		),
		Tab2: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("", ""),
		),
		Tab3: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("", ""),
		),
		Tab4: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("", ""),
		),
		Tab5: key.NewBinding(
			key.WithKeys("5"),
			key.WithHelp("", ""),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Up,
			k.Down,
		},
		{
			k.SwitchListPanel,
		},
		{
			k.Help,
			k.Quit,
		},
	}
}
