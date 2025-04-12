package asciiphysics

import (
	"image"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fogleman/gg"
	"github.com/qeesung/image2ascii/convert"
)

type Vector struct {
	X, Y float64
}
type Circle struct {
	color        lipgloss.TerminalColor
	radius       float64
	position     Vector
	velocity     Vector
	acceleration Vector
}
type Canvas struct {
	circles        []Circle
	width, height  int
	asciiConverter *convert.ImageConverter
}

func NewCanvas(width, height int) Canvas {
	converter := convert.NewImageConverter()
	return Canvas{
		width:          width,
		height:         height,
		asciiConverter: converter,
	}
}

func (c Canvas) Init() tea.Cmd {
	return nil
}

func (c Canvas) View() string {
	i := image.NewRGBA(image.Rect(0, 0, c.width, c.height))
	ctx := gg.NewContextForImage(i)
	for _, circle := range c.circles {
		circle.draw(ctx)
	}
	return c.asciiConverter.Image2ASCIIString(ctx.Image(), &convert.DefaultOptions)
}

func (c *Canvas) AddCircle(circle Circle) {
	c.circles = append(c.circles, circle)
}

func (c Canvas) Update(msg tea.Msg) (Canvas, tea.Cmd) {
	return c, nil
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

func (c Circle) draw(ctx *gg.Context) {
	ctx.Push()
	defer ctx.Pop()
	ctx.SetColor(c.color)
	ctx.DrawCircle(c.position.X, c.position.Y, c.radius)
	ctx.Fill()
}
