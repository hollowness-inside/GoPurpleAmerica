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
	c.CountiesN = sc.ReadInt()
	sc.Scan()
	return c
}
