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

func (sc *Scanner) ScanCounty() *County {
	c := new(County)
	c.Bbox = sc.ScanBBox()
	c.SubcountiesN = sc.ScanInt()

	c.Subcounties = make([]Subcounty, c.SubcountiesN)
	for i := 0; i < c.SubcountiesN; i++ {
		sc.Scan()

		name := sc.ScanString()
		countyName := sc.ScanString()
		n := sc.ScanInt()

		points := make([]Point, n)
		for j := 0; j < n; j++ {
			points[j] = sc.ScanPoint()
		}

		c.Subcounties[i] = Subcounty{
			name,
			countyName,
			n,
			points,
		}
	}
	return c
}
