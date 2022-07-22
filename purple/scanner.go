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
	sc.Scan()
	xy := strings.Split(sc.Text(), "   ")
	x, _ := strconv.ParseFloat(xy[0], 64)
	y, _ := strconv.ParseFloat(xy[1], 64)

	return Point{x, y}
}

func (sc *Scanner) ScanBBox() BBox {
	p1 := sc.ScanPoint()
	p2 := sc.ScanPoint()

	return BBox{p1, p2}
}

func (sc *Scanner) ScanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
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
