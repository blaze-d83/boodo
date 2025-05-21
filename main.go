package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	appName      = "boodo"
	dataFileName = "tasks.json"
)

// Styles
var (
	headerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).Padding(1, 2)
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true)
	taskStyle     = lipgloss.NewStyle().Padding(0, 2)
	doneStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	cursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	selectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("236")).Foreground(lipgloss.Color("230")).Padding(0, 1)
	frameStyle    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Margin(1)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
)

// Task represents a single TODO.
type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Model holds the state of the app.
type model struct {
	tasks     []Task
	cursor    int
	inputting bool
	input     textinput.Model
	err       error
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initialModel() model {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		tasks = []Task{}
	}

	ti := textinput.New()
	ti.Placeholder = "Add new task..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 30

	return model{
		tasks:     tasks,
		cursor:    len(tasks) - 1,
		inputting: false,
		input:     ti,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.inputting {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if title := m.input.Value(); title != "" {
					m.tasks = append(m.tasks, Task{Title: title})
					m.cursor = len(m.tasks) - 1
					m.err = saveTasks(m.tasks)
				}
				m.input.Reset()
				m.input.Blur()
				m.inputting = false
			case "esc":
				m.input.Reset()
				m.input.Blur()
				m.inputting = false
				m.err = nil
			}
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			m.inputting = true
			m.input.Focus()
			m.err = nil
		case "d":
			if m.validCursor() {
				m.tasks[m.cursor].Done = !m.tasks[m.cursor].Done
				m.err = saveTasks(m.tasks)
			}
		case "x":
			if m.validCursor() {
				m.tasks = slices.Delete(m.tasks, m.cursor, m.cursor+1)
				if len(m.tasks) == 0 {
					m.cursor = -1
				} else if m.cursor >= len(m.tasks) {
					m.cursor = len(m.tasks) - 1
				}
				m.err = saveTasks(m.tasks)
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "q", "ctrl+c":
			if err := saveTasks(m.tasks); err != nil {
				m.err = err
				return m, nil
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var header, body, inputView, footer strings.Builder

	// Header
	header.WriteString(headerStyle.Render("📋  BooDo — Your Tasks"))

	// Body
	for i, t := range m.tasks {
		line := fmt.Sprintf("%s %s", statusIcon(t.Done), t.Title)
		if m.cursor == i {
			body.WriteString(selectedStyle.Render(line) + "\n")
		} else {
			body.WriteString(taskStyle.Render(line) + "\n")
		}
	}

	// Input
	if m.inputting {
		inputView.WriteString(lipgloss.JoinVertical(0,
			helpStyle.Render("Type a new task and press Enter (esc to cancel):"),
			m.input.View(),
		))
	}

	// Footer
	footer.WriteString(helpStyle.Render("[n] New  [d] Toggle  [x] Delete  [j/k] Navigate  [q] Quit"))

	// Error message
	var errorMsg string
	if m.err != nil {
		errorMsg = "\n" + errorStyle.Render("Error: "+m.err.Error())
	}

	// Compose view
	screen := lipgloss.JoinVertical(0,
		header.String(),
		body.String(),
		inputView.String(),
		footer.String(),
		errorMsg,
	)

	return frameStyle.Render(screen)
}

func (m model) validCursor() bool {
	return len(m.tasks) > 0 && m.cursor >= 0 && m.cursor < len(m.tasks)
}

func statusIcon(done bool) string {
	if done {
		return doneStyle.Render("✔")
	}
	return "○"
}

// Data persistence
func getDataPath() (string, error) {
	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(cfg, appName)
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(appDir, dataFileName), nil
}

func loadTasks() ([]Task, error) {
	path, err := getDataPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	var ts []Task
	if err := json.Unmarshal(data, &ts); err != nil {
		return nil, err
	}
	return ts, nil
}

func saveTasks(ts []Task) error {
	path, err := getDataPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
