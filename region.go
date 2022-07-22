package main

type Region struct {
	Bbox        BBox
	SubregionsN int
	Subregions  []Region
}
