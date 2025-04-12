package asciiphysics

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fogleman/gg"
)

type Circle struct {
	color        lipgloss.TerminalColor
	radius       float64
	position     Vector
	velocity     Vector
	acceleration Vector
}

func (c *Circle) SetColor(color lipgloss.TerminalColor) bool {
	c.color = color
	return true
}

func (c *Circle) SetRadius(radius float64) bool {
	c.radius = radius
	return true
}

func (c *Circle) SetAcceleration(v Vector) bool {
	c.acceleration = v
	return true
}

func (c *Circle) SetVelocity(v Vector) bool {
	c.velocity = v
	return true
}

func (c *Circle) SetPosition(v Vector) bool {
	c.position = v
	return true
}

func (c Circle) Tick() Circle {
	x := c.position.X + c.velocity.X
	y := c.position.Y + c.velocity.Y
	c.position = Vector{X: x, Y: y}

	dx := c.velocity.X + c.acceleration.X
	dy := c.velocity.Y + c.acceleration.Y
	c.velocity = Vector{X: dx, Y: dy}
	return c
}

func (c Circle) Draw(ctx *gg.Context) {
	ctx.Push()
	defer ctx.Pop()
	ctx.SetColor(c.color)
	ctx.DrawCircle(c.position.X, c.position.Y, c.radius)
	ctx.Fill()
}
