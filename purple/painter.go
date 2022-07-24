package purple

import (
	"github.com/llgcode/draw2d/draw2dsvg"
)

type Painter struct {
	purple *Purple
}

func NewPainter(purple *Purple) Painter {
	painter := Painter{}
	painter.purple = purple
	return painter
}

func (p *Painter) Draw() {
	counties, _ := p.purple.ParseCounties()

	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	for _, county := range counties {
		p.drawCounty(county, gc)
	}

	draw2dsvg.SaveToSvgFile("myfile.svg", svg)
}

func (p *Painter) drawCounty(county *County, gc *draw2dsvg.GraphicContext) {
	minX := county.Bbox.MinX()
	maxY := county.Bbox.MaxY()

	gc.SetStrokeColor(p.purple.strokeColor)
	gc.SetLineWidth(p.purple.strokeWidth)

	for _, subc := range county.Subcounties {
		gc.BeginPath()

		start := subc.Points[0]
		xs, ys := start.x-minX, maxY-start.y
		gc.MoveTo(xs, ys)

		for _, point := range subc.Points {
			x := point.x - minX
			y := maxY - point.y
			gc.LineTo(x, y)
		}

		gc.LineTo(xs, ys)
		gc.Close()
		gc.FillStroke()
	}
}
