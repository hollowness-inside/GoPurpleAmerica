package purple

import (
	"archive/zip"
	"fmt"
	"image/color"

	"github.com/llgcode/draw2d/draw2dsvg"
)

type Purple struct {
	region string
	year   int
	colors [3][3]int

	strokeWidth float64
	strokeColor color.RGBA

	dataArchive    *zip.ReadCloser
	regionsArchive *zip.ReadCloser

	scale      float64
	outputPath string
}

func (p *Purple) UseDefault() {
	p.region = "USA"
	p.scale = 10
	p.strokeWidth = 0.2
	p.strokeColor = color.RGBA{0, 0, 0, 255}
	p.colors = [3][3]int{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
	}
}

func (p *Purple) ParseRegion() *Region {
	var zipFile *zip.File
	for _, f := range p.regionsArchive.File {
		if f.Name == p.region+".txt" {
			zipFile = f
		}
	}

	f, _ := zipFile.Open()
	defer f.Close()

	s := NewScanner(f)
	return s.ScanRegion()
}

func (p *Purple) Draw() {
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	region := p.ParseRegion()
	width := region.Bbox.MaxX() - region.Bbox.MinX()
	height := region.Bbox.MaxY() - region.Bbox.MinY()

	svg.Width = fmt.Sprintf("%fpx", width*p.scale)
	svg.Height = fmt.Sprintf("%fpx", height*p.scale)

	gc.Scale(p.scale, p.scale)
	p.drawRegion(region, gc)

	draw2dsvg.SaveToSvgFile(p.outputPath, svg)
}

func (p *Purple) drawRegion(region *Region, gc *draw2dsvg.GraphicContext) {
	minX := region.Bbox.MinX()
	maxY := region.Bbox.MaxY()

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.strokeWidth)

	for _, subc := range region.Subregions {
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
