package purple

type Region struct {
	Bbox        BBox
	Name        string
	SubregionsN int
	Subregions  []Subregion
}

type Subregion struct {
	Name       string
	RegionName string
	PointsN    int
	Points     []Point
}
