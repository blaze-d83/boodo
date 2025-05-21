package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	doneStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type Task struct {
	Title string
	Done  bool
}

type model struct {
	tasks     []Task
	cursor    int
	input     *textInput
	inputting bool
	lastKey   string
}

func NewModel() model {
	ti := &textInput{
		placeholder: "Enter task title",
		focus:       true,
	}

	return model{
		tasks:     []Task{},
		input:     ti,
		inputting: false,
	}
}

type textInput struct {
	text        string
	placeholder string
	focus       bool
}

func (ti *textInput) Init() tea.Cmd { return nil }

func (ti *textInput) View() string {
	if ti.text == "" {
		return fmt.Sprintf("[%s]", ti.placeholder)
	}
	return fmt.Sprintf("[%s]", ti.text)
}

func (ti *textInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return ti, nil
		case "esc":
			ti.text = ""
			return ti, nil
		case "backspace":
			if len(ti.text) > 0 {
				ti.text = ti.text[:len(ti.text)-1]
			}
		default:
			if len(msg.String()) == 1 {
				ti.text += msg.String()
			}
		}
	}
	return ti, nil
}

func (ti *textInput) Reset() { ti.text = "" }

func (ti *textInput) Value() string { return ti.text }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputting {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				title := strings.TrimSpace(m.input.Value())
				if title != "" {
					m.tasks = append(m.tasks, Task{Title: title})
					m.cursor = len(m.tasks) - 1
				}
				m.inputting = false
				m.input.Reset()
				return m, nil
			case "esc":
				m.inputting = false
				m.input.Reset()
				return m, nil
			}
		}

		var cmd tea.Cmd
		_, cmd = m.input.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "n":
			m.inputting = true
			return m, nil
		case "d":
			if len(m.tasks) > 0 {
				m.tasks[m.cursor].Done = !m.tasks[m.cursor].Done
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "x":
			if len(m.tasks) > 0 {
				m.tasks = slices.Delete(m.tasks, m.cursor, m.cursor+1)
				if m.cursor >= len(m.tasks) && len(m.tasks) > 0 {
					m.cursor = len(m.tasks) - 1
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.inputting {
		return fmt.Sprintf(
			"%s\n\n%s",
			helpStyle.Render("Add new task:"),
			m.input.View()+"\n\n"+helpStyle.Render("Enter to save . Esc to cancel"),
		)
	}

	var b strings.Builder
	b.WriteString("Your TODO List\n\n")

	for i, t := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		status := " "
		if t.Done {
			status = doneStyle.Render("✔ ")
		}

		title := t.Title
		if m.cursor == i {
			title = lipgloss.NewStyle().Bold(true).Render(title)
		}
		b.WriteString(fmt.Sprintf("%s %s %s\n", cursor, status, title))
	}

	help := helpStyle.Render("\n[n] New  .  [d] Toggle  . [x] Delete  .  [Up/Down or k/j] Navigate  .  [q] Quit\n")
	b.WriteString(help)

	return b.String()
}

func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
