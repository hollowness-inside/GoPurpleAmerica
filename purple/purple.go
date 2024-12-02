package purple

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dsvg"
)

type RGBA = color.RGBA

type Purple struct {
	Region *Region
	Year   string

	Stats map[string]RGBA

	Scale       float64
	StrokeWidth float64
	StrokeColor RGBA
}

func (p *Purple) GenerateSVG() *draw2dsvg.Svg {
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	width := p.Region.Bbox.Max.X - p.Region.Bbox.Min.X
	height := p.Region.Bbox.Max.Y - p.Region.Bbox.Min.Y

	svg.Width = fmt.Sprintf("%fpx", width*p.Scale)
	svg.Height = fmt.Sprintf("%fpx", height*p.Scale)

	gc.Scale(p.Scale, p.Scale)
	p.draw(gc)

	return svg
}

func (p *Purple) draw(gc *draw2dsvg.GraphicContext) {
	gc.SetStrokeColor(p.StrokeColor)
	gc.SetLineWidth(p.StrokeWidth)

	subregions := p.outlineSubregions()
	for _, subregion := range subregions {
		p.drawSubregion(gc, &subregion)
	}
}

func (p *Purple) drawSubregion(gc *draw2dsvg.GraphicContext, subregion *Subregion) {
	subregionColor := p.getSubregionColor(subregion.Name)
	gc.SetFillColor(subregionColor)
	gc.Fill(subregion.Path)
	gc.Stroke(subregion.Path)
}

// Extract a color for a given county from the statistics
func (p *Purple) getSubregionColor(subregion string) RGBA {
	if v, ok := p.Stats[subregion]; ok {
		return v
	}
	return RGBA{0, 0, 0, 0}
}

// Concurrently outline all counties and get their paths
func (p *Purple) outlineSubregions() []Subregion {
	wg := new(sync.WaitGroup)
	wg.Add(len(p.Region.Subregions))
	for i := range p.Region.Subregions {
		go p.outlineSubregion(wg, &p.Region.Subregions[i])
	}
	wg.Wait()
	return p.Region.Subregions
}

// Draws a county outline and sends it to the counties channel
func (p *Purple) outlineSubregion(wg *sync.WaitGroup, subregion *Subregion) {
	defer wg.Done()

	pointsN := len(subregion.Points)

	path := new(draw2d.Path)
	path.Components = make([]draw2d.PathCmp, pointsN+1)
	path.Points = make([]float64, 2*pointsN+2)

	for i, point := range subregion.Points {
		x := point.X - p.Region.Bbox.Min.X
		y := p.Region.Bbox.Max.Y - point.Y

		path.Components[i] = draw2d.LineToCmp
		path.Points[i*2] = x
		path.Points[i*2+1] = y
	}

	path.Components[0] = draw2d.MoveToCmp
	path.Components[pointsN] = draw2d.LineToCmp
	path.Points[2*pointsN] = path.Points[0]
	path.Points[2*pointsN+1] = path.Points[1]

	subregion.Path = path
}
