package main

import (
	"fmt"
	"math/rand"

	"github.com/Andrew-Wichmann/fireworks-cli/pkg/asciiphysics"
	"github.com/Andrew-Wichmann/fireworks-cli/pkg/firework"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const INIT state = 0
const RUNNING state = 1

type model struct {
	state     state
	canvas    asciiphysics.Canvas
	fireworks []firework.Model
}

func newModel(width, height int) model {
	canvas := asciiphysics.NewCanvas(width, height)
	for range 10 {
		x := float64(width) * rand.Float64()
		y := float64(height) * rand.Float64()
		loc := asciiphysics.Vector{x, y}
		f := firework.New(firework.ShortFuse, firework.ShortFuse, loc, firework.RandomColor())
		canvas.AddDrawable(f)
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
