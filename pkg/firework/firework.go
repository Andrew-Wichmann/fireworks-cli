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

type charge []asciiphysics.Circle

type chargeOption func(*charge)

var defaultColor fireworkColor = Red

const defaultRadius float64 = 2
const defaultBurstChargeSize float64 = 2

func WithColor(color fireworkColor) chargeOption {
	return func(c *charge) {
		for i, star := range *c {
			star.SetColor(color.sequence[0])
			(*c)[i] = star
		}
	}
}
func WithRadius(radius float64) chargeOption {
	return func(c *charge) {
		for i, star := range *c {
			star.SetRadius(radius)
			(*c)[i] = star
		}
	}
}
func WithBurstChargeSize(burstChargeSize float64) chargeOption {
	return func(c *charge) {
		angle := 0.0
		for i, star := range *c {
			dx := math.Cos(degToRad(angle)) * burstChargeSize
			dy := math.Sin(degToRad(angle)) * burstChargeSize
			star.SetVelocity(asciiphysics.Vector{X: dx, Y: dy})
			star.SetAcceleration(asciiphysics.Vector{X: 0.0, Y: .1})
			angle += float64(i) * float64((360 / len(*c)))
			(*c)[i] = star
		}
	}
}
func New(start asciiphysics.Vector, chargeOptions ...chargeOption) Model {
	charges := make(charge, 50)
	angle := 0.0
	for i, star := range charges {
		star.SetRadius(defaultRadius)
		dx := math.Cos(degToRad(angle)) * defaultBurstChargeSize
		dy := math.Sin(degToRad(angle)) * defaultBurstChargeSize
		star.SetVelocity(asciiphysics.Vector{X: dx, Y: dy})
		star.SetAcceleration(asciiphysics.Vector{X: 0.0, Y: .1})
		star.SetPosition(start)
		angle += float64(i) * float64((360 / len(charges)))
		star.SetColor(Red.sequence[0])
		charges[i] = star
	}
	for _, opt := range chargeOptions {
		opt(&charges)
	}
	return Model{charge: charges}
}

type Model struct {
	state  state
	charge charge
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
