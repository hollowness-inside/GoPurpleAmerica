package purple

import (
	"archive/zip"
	"fmt"
	"image/color"

	"github.com/llgcode/draw2d/draw2dsvg"
)

type Purple struct {
	county string
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
	p.county = "USA"
	p.scale = 10
	p.strokeWidth = 0.2
	p.strokeColor = color.RGBA{0, 0, 0, 255}
	p.colors = [3][3]int{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
	}
}

func (p *Purple) ParseCounty() *County {
	var zipFile *zip.File
	for _, f := range p.regionsArchive.File {
		if f.Name == p.county+".txt" {
			zipFile = f
		}
	}

	f, _ := zipFile.Open()
	defer f.Close()

	s := NewScanner(f)
	return s.ScanCounty()
}

func (p *Purple) Draw() {
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	county := p.ParseCounty()
	width := county.Bbox.MaxX() - county.Bbox.MinX()
	height := county.Bbox.MaxY() - county.Bbox.MinY()

	svg.Width = fmt.Sprintf("%fpx", width*p.scale)
	svg.Height = fmt.Sprintf("%fpx", height*p.scale)

	gc.Scale(p.scale, p.scale)
	p.drawCounty(county, gc)

	draw2dsvg.SaveToSvgFile(p.outputPath, svg)
}

func (p *Purple) drawCounty(county *County, gc *draw2dsvg.GraphicContext) {
	minX := county.Bbox.MinX()
	maxY := county.Bbox.MaxY()

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.strokeWidth)

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
