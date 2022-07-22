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

func NewScanner(r io.Reader) Scanner {
	return Scanner{*bufio.NewScanner(r)}
}

func (sc *Scanner) ReadPoint() Point {
	sc.Scan()
	xy := strings.Split(sc.Text(), "   ")
	x, _ := strconv.ParseFloat(xy[0], 64)
	y, _ := strconv.ParseFloat(xy[1], 64)

	return Point{x, y}
}
