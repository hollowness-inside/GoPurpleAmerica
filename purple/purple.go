package purple

import (
	"archive/zip"
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
}

func (p *Purple) UseDefault() {
	p.county = "USA"
	p.strokeWidth = 0.5
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
	county := p.ParseCounty()

	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	p.drawCounty(county, gc)

	draw2dsvg.SaveToSvgFile("myfile.svg", svg)
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
