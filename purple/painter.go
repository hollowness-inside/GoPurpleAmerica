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

	img := image.NewGray(image.Rect(x0, y0, x1, y1))

	subc := county.Subcounties[1]
	for _, point := range subc.Points {
		x := int(point.x * scale)
		y := int(point.y * scale)
		img.Set(x, y, color.White)
	}

	f, _ := os.Create("myfile.png")
	defer f.Close()
	png.Encode(f, img)
}
