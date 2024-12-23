package purple

import "github.com/llgcode/draw2d"

// Region represents a geographical area and its constituent subregions.
// It can represent any hierarchical level like a country containing regions,
// a region containing subregions, or any other administrative division.
type Region struct {
	Bbox       BBox        // Bounding box containing the entire region
	Name       string      // Name of the region
	Subregions []Subregion // Slice containing all subregions
}

// Subregion represents a geographical subdivision within a parent region.
// It contains the subregion's name, its parent region's name, the points
// that define its boundary, and the path representing its boundary.
type Subregion struct {
	Name       string       // Name of the subregion
	ParentName string       // Name of the parent region this subregion belongs to
	Points     []Point      // Slice of points defining the subregion boundary
	Path       *draw2d.Path // Draw2d path representing the subregion boundary
}
