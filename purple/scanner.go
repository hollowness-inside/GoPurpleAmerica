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

func (sc *Scanner) ReadPoint() Point {
	sc.Scan()
	xy := strings.Split(sc.Text(), "   ")
	x, _ := strconv.ParseFloat(xy[0], 64)
	y, _ := strconv.ParseFloat(xy[1], 64)

	return Point{x, y}
}

func (sc *Scanner) ReadBBox() BBox {
	p1 := sc.ReadPoint()
	p2 := sc.ReadPoint()

	return BBox{p1, p2}
}

func (sc *Scanner) ReadInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (sc *Scanner) ReadCounty() *County {
	c := new(County)
	c.Bbox = sc.ReadBBox()
	c.SubcountiesN = sc.ReadInt()

	c.Subcounties = make([]Subcounty, c.SubcountiesN)
	for i := 0; i < c.SubcountiesN; i++ {
		sc.Scan()

		name := sc.ReadString()
		countyName := sc.ReadString()
		n := sc.ReadInt()

		points := make([]Point, n)
		for j := 0; j < n; j++ {
			points[j] = sc.ReadPoint()
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

func (sc *Scanner) ReadString() string {
	sc.Scan()
	return sc.Text()
}
