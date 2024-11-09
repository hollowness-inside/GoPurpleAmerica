package purple

type Point struct {
	X float64
	Y float64
}

type BBox struct {
	Point1 Point
	Point2 Point

	Max Point
	Min Point
}

func NewBBox(p1, p2 Point) BBox {
	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	if x1 < x2 {
		x1, x2 = x2, x1
	}

	if y1 < y2 {
		y1, y2 = y2, y1
	}

	maxX, minX := x1, x2
	maxY, minY := y1, y2

	return BBox{
		Point1: p1,
		Point2: p2,
		Max:    Point{maxX, maxY},
		Min:    Point{minX, minY},
	}
}
