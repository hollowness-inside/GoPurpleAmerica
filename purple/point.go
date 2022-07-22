package purple

type Point struct {
	x, y float64
}

type BBox struct {
	p1, p2 Point
}

func (b *BBox) MaxX() float64 {
	if b.p1.x > b.p2.x {
		return b.p1.x
	} else {
		return b.p2.x
	}
}

func (b *BBox) MaxY() float64 {
	if b.p1.y > b.p2.y {
		return b.p1.y
	} else {
		return b.p2.y
	}
}
