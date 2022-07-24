package purple

import (
	"archive/zip"
	"image/color"
	"sync"

	"github.com/llgcode/draw2d/draw2dsvg"
)

type Purple struct {
	counties []string
	year     int
	colors   [3][3]int

	strokeWidth float64
	strokeColor color.RGBA

	dataArchive    *zip.ReadCloser
	regionsArchive *zip.ReadCloser
}

func (p *Purple) ParseCounties() ([]*County, error) {
	files := make(map[string]*zip.File, 0)
	for _, county := range p.counties {
		files[county+".txt"] = nil
	}

	for _, v := range p.regionsArchive.File {
		if _, ok := files[v.Name]; ok {
			files[v.Name] = v
		}
	}

	if len(files) != len(p.counties) {
		return nil, ErrCountyName
	}

	wg := &sync.WaitGroup{}
	ch := make(chan *County, len(p.counties))

	for _, file := range files {
		wg.Add(1)
		go fileReader(wg, ch, file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	counties := make([]*County, len(p.counties))

	i := 0
	for c := range ch {
		counties[i] = c
		i++
	}

	return counties, nil
}

func fileReader(wg *sync.WaitGroup, counties chan *County, zipFile *zip.File) {
	defer wg.Done()

	f, _ := zipFile.Open()
	defer f.Close()

	s := NewScanner(f)
	counties <- s.ScanCounty()
}

func (p *Purple) Draw() {
	counties, _ := p.ParseCounties()

	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)

	for _, county := range counties {
		p.drawCounty(county, gc)
	}

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
