package purple

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Painter struct {
	purple *Purple
	image  *image.RGBA
}

func NewPainter(purple *Purple) Painter {
	painter := Painter{}
	painter.purple = purple
	return painter
}

func (p *Painter) Draw() {
	counties, _ := p.purple.ParseCounties()
	county := counties[0]

	scale := float64(p.purple.scale)

	x0 := county.Bbox.MinX() * scale
	y0 := county.Bbox.MinY() * scale

	x1 := county.Bbox.MaxX() * scale
	y1 := county.Bbox.MaxY() * scale

	width := int(x1 - x0)
	height := int(y1 - y0)

	img := image.NewRGBA(image.Rect(0, 0, width+1, height+1))
	p.image = img

	for _, subc := range county.Subcounties {
		points := make([]Point, subc.PointsN)
		for i, point := range subc.Points {
			x := point.x*scale - x0
			y := y1 - point.y*scale
			// img.Set(int(x), int(y), color.White)
			points[i] = Point{x, y}
		}
		p.DrawOutline(&points)
	}

	f, _ := os.Create("myfile.png")
	defer f.Close()

	png.Encode(f, img)
}

func (p *Painter) DrawOutline(pts *[]Point) {
	points := *pts
	prev := points[0]
	for _, pt := range points[1:] {
		p.DrawLine(pt, prev)
		prev = pt
	}

	p.DrawLine(points[0], points[len(points)-1])
}

func (p *Painter) DrawLine(start, end Point) {
	x0, y0 := start.x, start.y
	x1, y1 := end.x, end.y

	if x0 == x1 {
		if y1 > y0 {
			y0, y1 = y1, y0
		}

		x := int(x0)
		for y := y0; y <= y1; y += 0.1 {
			p.image.Set(x, int(y), color.White)
		}
		return
	}

	if y0 == y1 {
		if x1 > x0 {
			x0, x1 = x1, x0
		}

		y := int(y0)
		for x := x0; x <= x1; x += 0.1 {
			p.image.Set(int(x), y, color.White)
		}
		return
	}

	dx := x1 - x0
	dy := y1 - y0

	distance := math.Sqrt(dx*dx + dy*dy)

	angle := math.Atan2(dy, dx)

	for i := 0.0; i < distance; i += 0.01 {
		x := x0 + math.Cos(angle)*i
		y := y0 + math.Sin(angle)*i

		p.image.Set(int(x), int(y), color.White)
	}
}
