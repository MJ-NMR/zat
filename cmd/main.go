package main

import (
	"fmt"
	"log"
	//"math/rand"
	"zat/backend"

	tea "github.com/charmbracelet/bubbletea"
)

type location struct {
	y int
	x int
}

type model struct {
	status    backend.Status
	locations map[string]location
	width     int
	height    int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.status = backend.RandomStatus(m.height, m.width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m model) View() (s string) {
	for y := range len(m.status) {
		for x := range len(m.status[0]) {
			s += fmt.Sprintf("%T ", m.status[y][x])
		}
		s += "\n"
	}
	return s
}

func initModel() tea.Model {
	m := model{}
	m.locations = make(map[string]location)
	return m
}

func main() {
	app := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		log.Panic(err)
	}
}
