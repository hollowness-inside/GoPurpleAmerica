package purple

// Point represents a 2D coordinate with X and Y values
type Point struct {
	X float64 // X is the x-coordinate of the point
	Y float64 // Y is the y-coordinate of the point
}

// BBox represents a bounding box defined by two points.
type BBox struct {
	Max Point // Upper right point
	Min Point // Lower left point
}

// NewBBox creates a new bounding box from two points.
func NewBBox(p1, p2 Point) BBox {
	return BBox{
		Max: Point{
			X: max(p1.X, p2.X),
			Y: max(p1.Y, p2.Y),
		},
		Min: Point{
			X: min(p1.X, p2.X),
			Y: min(p1.Y, p2.Y),
		},
	}
}
