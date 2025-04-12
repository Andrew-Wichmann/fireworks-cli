package firework

import (
	"math"

	"github.com/Andrew-Wichmann/fireworks-cli/pkg/asciiphysics"
	"github.com/fogleman/gg"
)

type state int

const (
	lighting  state = 0
	flying    state = 1
	exploding state = 2
)

type fuseLength int

const (
	ShortFuse  = 0
	MediumFuse = 1
	LongFuse   = 2
)

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func New(initFuse, flightFuse fuseLength, start asciiphysics.Vector, color fireworkColor) Model {
	charge := make([]asciiphysics.Circle, 50)
	angle := 0.0
	for i, particle := range charge {
		particle.SetColor(color.sequence[0])
		particle.SetRadius(2)
		dx := math.Cos(degToRad(angle))
		dy := math.Sin(degToRad(angle))
		particle.SetVelocity(asciiphysics.Vector{X: dx, Y: dy})
		particle.SetAcceleration(asciiphysics.Vector{X: 0.0, Y: .1})
		particle.SetPosition(start)
		charge[i] = particle
		angle += float64(i) * float64((360 / len(charge)))
	}

	return Model{charge: charge}
}

type Model struct {
	color      fireworkColor
	state      state
	charge     []asciiphysics.Circle
	initFuse   int
	flightFuse int
}

func (m Model) Tick() asciiphysics.Drawable {
	for i, particle := range m.charge {
		m.charge[i] = particle.Tick()
	}
	return m
}

func (m Model) Draw(ctx *gg.Context) {
	for _, particle := range m.charge {
		particle.Draw(ctx)
	}
}
