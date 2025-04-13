package firework

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type fireworkColor struct {
	sequence []lipgloss.TerminalColor
}

var Green = fireworkColor{
	sequence: []lipgloss.TerminalColor{lipgloss.Color("#009900"), lipgloss.Color("#006600"), lipgloss.Color("#006622"), lipgloss.Color("#00ff00"), lipgloss.Color("#00ff33"), lipgloss.Color("#008833")},
}
var Red = fireworkColor{
	sequence: []lipgloss.TerminalColor{lipgloss.Color("#990000"), lipgloss.Color("#660000"), lipgloss.Color("#662200"), lipgloss.Color("#ff0000"), lipgloss.Color("#ff3300"), lipgloss.Color("#883300")},
}
var Blue = fireworkColor{
	sequence: []lipgloss.TerminalColor{lipgloss.Color("#000099"), lipgloss.Color("#000066"), lipgloss.Color("#220066"), lipgloss.Color("#0000ff"), lipgloss.Color("#3300ff"), lipgloss.Color("#330088")},
}

var allColors = []fireworkColor{Blue, Red, Green}

func RandomColor() fireworkColor {
	i := rand.Intn(len(allColors))
	return allColors[i]
}
