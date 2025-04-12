package main

import (
	"fmt"
	"math/rand"

	"github.com/Andrew-Wichmann/fireworks-cli/pkg/asciiphysics"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const INIT state = 0
const RUNNING state = 1

type model struct {
	state  state
	canvas asciiphysics.Canvas
}

var colors []lipgloss.TerminalColor = []lipgloss.TerminalColor{lipgloss.Color("#990000"), lipgloss.Color("#660000"), lipgloss.Color("#662200"), lipgloss.Color("#ff0000"), lipgloss.Color("#ff3300"), lipgloss.Color("#883300")}

func newModel(width, height int) model {
	canvas := asciiphysics.NewCanvas(width, height)
	particles := make([]asciiphysics.Circle, 50)
	for _, particle := range particles {
		color := colors[rand.Intn(len(colors))]
		particle.SetColor(color)
		particle.SetRadius(5)
		x := rand.Float64() * float64(rand.Intn(width))
		y := rand.Float64() * float64(rand.Intn(height))
		particle.SetPosition(asciiphysics.Vector{X: x, Y: y})
		canvas.AddCircle(particle)
	}
	return model{
		canvas: canvas,
	}
}

func (m model) Init() tea.Cmd {
	return m.canvas.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m := newModel(msg.Width, msg.Height*2)
		m.state = RUNNING
		return m, nil
	}
	if m.state == RUNNING {
		canvas, cmd := m.canvas.Update(msg)
		m.canvas = canvas
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.state == RUNNING {
		return m.canvas.View()
	} else {
		return fmt.Sprintf("%d", m.state)
	}
}
