package asciiphysics

import (
	"image"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/gg"
	"github.com/qeesung/image2ascii/convert"
)

type canvasTick struct{}

func newTick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg {
		return canvasTick{}
	})
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
	return newTick()
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
	if _, ok := msg.(canvasTick); ok {
		for i, circle := range c.circles {
			c.circles[i] = circle.update()
		}
		return c, newTick()
	}
	return c, nil
}
