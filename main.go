package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	Title string
	Done  bool
}

type model struct {
	tasks  []Task
	cursor int
}

func NewModel() model {
	return model{
		tasks:  []Task{},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "n":
		// TODO: Implement adding a new task
		case "d":
		// TODO: Toggle m.tasks[m.cursor].Done
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Your TODO list:\n\n"
	for i, t := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		status := " "
		if t.Done {
			status = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, status, t.Title)
	}
	s += "\n[n] Add new task     [d] Toggle done     [q] Quit\n"
	return s
}

func main() {
	p := tea.NewProgram(NewModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running the program", err)
		os.Exit(1)
	}

}
