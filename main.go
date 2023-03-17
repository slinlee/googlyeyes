package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type errMsg error

type model struct {
	// spinner      spinner.Model
	err          error
	x            int
	y            int
	canvasWidth  int
	canvasHeight int
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "press q to quit"),
)

func initialModel() model {
	// s := spinner.New()
	// s.Spinner = spinner.Dot
	// s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	// return model{spinner: s}
	return model{}
}

func (m model) Init() tea.Cmd {
	// return m.spinner.Tick
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			return m, tea.Quit

		}
		return m, nil
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			m.x = msg.X
			m.y = msg.Y

			return m, nil
		}
	case tea.WindowSizeMsg:
		// set window size for calculations
		m.canvasWidth, m.canvasHeight = msg.Width, msg.Height
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		// m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\nMouse: %d %d\nWindow: %d %d",
		// m.spinner.View(),
		m.x,
		m.y,
		m.canvasWidth,
		m.canvasHeight,)

	return str
}

func main() {
	p := tea.NewProgram(initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
