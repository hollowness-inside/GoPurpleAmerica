package purple

import (
	"fmt"
	"image/color"

	"github.com/llgcode/draw2d/draw2dsvg"
)

type RGBA = color.RGBA

type Purple struct {
	State *State
	Year  string

	Stats      map[string]RGBA
	OutputPath string

	Scale       float64
	StrokeWidth float64
	StrokeColor RGBA
}

func (p *Purple) UseDefaults() {
	p.Scale = 10
	p.StrokeWidth = 0.2
	p.StrokeColor = RGBA{0, 0, 0, 255}
}

func (p *Purple) Draw() {
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	width := p.State.Bbox.Max.X - p.State.Bbox.Min.X
	height := p.State.Bbox.Max.Y - p.State.Bbox.Min.Y

	svg.Width = fmt.Sprintf("%fpx", width*p.Scale)
	svg.Height = fmt.Sprintf("%fpx", height*p.Scale)

	gc.Scale(p.Scale, p.Scale)
	p.drawState(p.State, gc)

	draw2dsvg.SaveToSvgFile(p.OutputPath, svg)
}

func (p *Purple) drawState(state *State, gc *draw2dsvg.GraphicContext) {
	gc.SetStrokeColor(p.StrokeColor)
	gc.SetLineWidth(p.StrokeWidth)

	for _, county := range state.Counties {
		clr := p.getCountyColor(county.Name)
		gc.SetFillColor(clr)
		gc.BeginPath()

		start := county.Points[0]
		xs := start.X - state.Bbox.Min.X
		ys := state.Bbox.Max.Y - start.Y
		gc.MoveTo(xs, ys)

		for _, point := range county.Points {
			x := point.X - state.Bbox.Min.X
			y := state.Bbox.Max.Y - point.Y
			gc.LineTo(x, y)
		}

		gc.LineTo(xs, ys)
		gc.Close()
		gc.FillStroke()
		gc.Fill(gc.Current.Path)
	}
}

func (p *Purple) getCountyColor(county string) RGBA {
	if v, ok := p.Stats[county]; ok {
		return v
	}
	return RGBA{0, 0, 0, 0}
}
