package main

import (
	"log"
	"zat/backend"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	frame backend.Frame
	name   string
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Zzz"
	}

	return "hello"
}

func main() {
	app := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		log.Panic(err)
	}
}
