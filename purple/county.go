package purple

import "github.com/llgcode/draw2d"

type State struct {
	Bbox      BBox
	Name      string
	CountiesN int
	Counties  []County
}

type County struct {
	Name      string
	StateName string
	PointsN   int
	Points    []Point
}

type ChanCounty struct {
	Name string
	Path *draw2d.Path
}
