package purple

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Scanner struct {
	bufio.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{*bufio.NewScanner(r)}
}

func (sc *Scanner) ScanPoint() Point {
	xy := strings.Split(sc.ScanString(), "   ")

	xs := strings.TrimSpace(xy[0])
	ys := strings.TrimSpace(xy[1])

	x, _ := strconv.ParseFloat(xs, 64)
	y, _ := strconv.ParseFloat(ys, 64)

	return Point{x, y}
}

func (sc *Scanner) ScanBBox() BBox {
	p1 := sc.ScanPoint()
	p2 := sc.ScanPoint()

	return BBox{p1, p2}
}

func (sc *Scanner) ScanInt() int {
	v, _ := strconv.Atoi(sc.ScanString())
	return v
}

func (sc *Scanner) ScanString() string {
	sc.Scan()
	return sc.Text()
}

func (sc *Scanner) ScanRegion() *Region {
	reg := new(Region)
	reg.Bbox = sc.ScanBBox()
	reg.SubregionsN = sc.ScanInt()

	reg.Subregions = make([]Subregion, reg.SubregionsN)
	for i := 0; i < reg.SubregionsN; i++ {
		sc.Scan()

		name := sc.ScanString()
		regionName := sc.ScanString()
		n := sc.ScanInt()

		points := make([]Point, n)
		for j := 0; j < n; j++ {
			points[j] = sc.ScanPoint()
		}

		reg.Subregions[i] = Subregion{
			name,
			regionName,
			n,
			points,
		}
	}
	return reg
}
