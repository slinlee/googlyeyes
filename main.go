package main

import (
	"fmt"
	"math"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type model struct {
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
	return model{}
}

func (m model) Init() tea.Cmd {
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
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\nMouse: %d %d\nWindow: %d %d\n\n%s",
		m.x,
		m.y,
		m.canvasWidth,
		m.canvasHeight,
		drawCircle(16),
	)

	return str
}

func drawCircle(radius int) string {
	s := "--\n"
	vChars := 2 * radius
	hChars := 2 * radius * 2      // double the width for each 'pixel'
	center := float64(radius + 2) // hard coded

	for j := 0; j <= vChars-1; j++ {
		y := float64(j) + float64(0.5)
		for i := 0; i <= hChars-1; i++ {
			// edge := float64(x*x+y*y) / float64(r-r)
			x := 2 * (float64(i) + float64(0.5))
			dist := math.Sqrt((x-center)*(x-center) + (y-center)*(y-center))

			if dist > float64(radius-1) && dist < float64(radius+1) {
				s += "**"
			} else {
				s += "--"
			}
		}
		s += "\n"
	}
	return s

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
