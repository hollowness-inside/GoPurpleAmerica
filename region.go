package main

type Region struct {
	Bbox        BBox
	SubregionsN int
	Subregions  []Subregion
}

type Subregion struct {
	PointsN    int
	Name       string
	RegionName string
}
