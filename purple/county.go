package purple

type County struct {
	Bbox         BBox
	Name         string
	SubcountiesN int
	Subcounties  []Subcounty
}

type Subcounty struct {
	Name       string
	CountyName string
	PointsN    int
	Points     []Point
}
