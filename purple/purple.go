package purple

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/llgcode/draw2d/draw2dsvg"
)

type Purple struct {
	regionName string
	year       string

	strokeWidth float64
	strokeColor color.RGBA

	region *Region
	stats  map[string]color.RGBA

	scale      float64
	outputPath string
}

func (p *Purple) UseDefault() {
	p.regionName = "USA"
	p.scale = 10
	p.strokeWidth = 0.2
	p.strokeColor = color.RGBA{0, 0, 0, 255}
}

func (p *Purple) Draw() {
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	width := p.region.Bbox.Max.X - p.region.Bbox.Min.X
	height := p.region.Bbox.Max.Y - p.region.Bbox.Min.Y

	svg.Width = fmt.Sprintf("%fpx", width*p.scale)
	svg.Height = fmt.Sprintf("%fpx", height*p.scale)

	gc.Scale(p.scale, p.scale)
	p.drawRegion(p.region, gc)

	draw2dsvg.SaveToSvgFile(p.outputPath, svg)
}

func (p *Purple) drawRegion(region *Region, gc *draw2dsvg.GraphicContext) {
	minX := region.Bbox.Min.Y
	maxY := region.Bbox.Max.Y

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.strokeWidth)

	for _, subc := range region.Subregions {
		gc.SetFillColor(p.GetSubregionColor(subc.Name))
		gc.BeginPath()

		start := subc.Points[0]
		xs, ys := start.X-minX, maxY-start.Y
		gc.MoveTo(xs, ys)

		for _, point := range subc.Points {
			x := point.X - minX
			y := maxY - point.Y
			gc.LineTo(x, y)
		}

		gc.LineTo(xs, ys)
		gc.Close()
		gc.FillStroke()
		gc.Fill(gc.Current.Path)
	}
}

func ReadStatistics(r io.Reader) map[string]color.RGBA {
	data := make(map[string]color.RGBA, 0)

	reader := bufio.NewScanner(r)

	// Skipping the header
	reader.Scan()

	for reader.Scan() {
		row := strings.Split(reader.Text(), ",")

		r1, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil
		}

		r2, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil
		}

		r3, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil
		}

		sum := r1 + r2 + r3

		r := uint8((r1 / sum) * 255)
		g := uint8((r2 / sum) * 255)
		b := uint8((r3 / sum) * 255)

		data[row[0]] = color.RGBA{r, g, b, 255}
	}

	return data
}

func (p *Purple) GetSubregionColor(subregion string) color.RGBA {
	if v, ok := p.stats[subregion]; ok {
		return v
	}
	return color.RGBA{0, 0, 0, 0}
}
