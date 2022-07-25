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
	p.drawState(gc)

	draw2dsvg.SaveToSvgFile(p.OutputPath, svg)
}

func (p *Purple) drawState(gc *draw2dsvg.GraphicContext) {
	gc.SetStrokeColor(p.StrokeColor)
	gc.SetLineWidth(p.StrokeWidth)

	counties := p.projectCounties()

	for _, county := range counties {
		clr := p.getCountyColor(county.Name)
		gc.SetFillColor(clr)
		gc.BeginPath()

		start := county.Points[0]

		gc.MoveTo(start.X, start.Y)
		for _, point := range county.Points {
			gc.LineTo(point.X, point.Y)
		}
		gc.LineTo(start.X, start.Y)

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

func (p *Purple) projectCounties() []County {
	counties := make([]County, p.State.CountiesN)

	for n, county := range p.State.Counties {
		newCounty := County{}

		points := make([]Point, county.PointsN)
		for i, point := range county.Points {
			x := point.X - p.State.Bbox.Min.X
			y := p.State.Bbox.Max.Y - point.Y
			points[i] = Point{x, y}
		}

		newCounty.Name = county.Name
		newCounty.PointsN = county.PointsN
		newCounty.Points = points

		counties[n] = newCounty
	}

	return counties
}
