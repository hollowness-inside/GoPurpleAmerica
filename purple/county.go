package purple

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
