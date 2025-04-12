package asciiphysics

import (
	"image"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/gg"
	"github.com/qeesung/image2ascii/convert"
)

const (
	fps = 60
)

type canvasTick struct{}
type Drawable interface {
	Tick() Drawable
	Draw(*gg.Context)
}

func newTick() tea.Cmd {
	return tea.Tick(time.Second/fps, func(time.Time) tea.Msg {
		return canvasTick{}
	})
}

type Canvas struct {
	drawable       []Drawable
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
	for _, drawable := range c.drawable {
		drawable.Draw(ctx)
	}
	return c.asciiConverter.Image2ASCIIString(ctx.Image(), &convert.DefaultOptions)
}

func (c *Canvas) AddDrawable(drawable Drawable) {
	c.drawable = append(c.drawable, drawable)
}

func (c Canvas) Update(msg tea.Msg) (Canvas, tea.Cmd) {
	if _, ok := msg.(canvasTick); ok {
		for i, circle := range c.drawable {
			c.drawable[i] = circle.Tick()
		}
		return c, newTick()
	}
	return c, nil
}
