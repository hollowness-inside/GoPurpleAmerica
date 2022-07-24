package purple

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Painter struct {
	purple *Purple
}

func NewPainter(p *Purple) Painter {
	return Painter{p}
}

func (p *Painter) Draw() {
	counties, _ := p.purple.ParseCounties()
	county := counties[0]

	scale := float64(p.purple.scale)

	x0 := int(county.Bbox.MinX() * scale)
	y0 := int(county.Bbox.MinY() * scale)

	x1 := int(county.Bbox.MaxX() * scale)
	y1 := int(county.Bbox.MaxY() * scale)

	width := x1 - x0
	height := y1 - y0

	img := image.NewGray(image.Rect(0, 0, width, height))

	for _, subc := range county.Subcounties {
		for _, point := range subc.Points {
			x := x1 - int(point.x*scale)
			y := int(point.y*scale) - y0
			img.Set(x, y, color.White)
		}
	}

	f, _ := os.Create("myfile.png")
	defer f.Close()

	png.Encode(f, img)
}

func (p *Painter) DrawOutline(pts *[]Point) {
	points := *pts

	prev := points[0]
	for _, p := range points[1:] {
		p.DrawLine(p, prev)
		prev = p
	}

	p.DrawLine(points[0], points[len(points)-1])
}
