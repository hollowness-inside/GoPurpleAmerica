package purple

import "github.com/llgcode/draw2d"

// Region represents a geographical area and its constituent subregions.
// It can represent any hierarchical level like a country containing states,
// a state containing counties, or any other administrative division.
type Region struct {
	Bbox       BBox        // Bounding box containing the entire region
	Name       string      // Name of the region
	Subregions []Subregion // Slice containing all subregions
}

// Subregion represents a geographical subdivision within a parent region.
// It contains the subregion's name, its parent region's name, and the points
// that define its boundary.
type Subregion struct {
	Name       string  // Name of the subregion
	ParentName string  // Name of the parent region this subregion belongs to
	Points     []Point // Slice of points defining the subregion boundary
}

// TODO: Embed into Subregion

// RegionPath represents a region's boundary as a drawable path.
// It associates a region's name with its graphical representation.
type RegionPath struct {
	Name string       // Name of the region
	Path *draw2d.Path // Draw2d path representing the region boundary
}
