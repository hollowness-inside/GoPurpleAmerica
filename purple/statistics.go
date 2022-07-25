package purple

import (
	"bufio"
	"image/color"
	"io"
	"strconv"
	"strings"
)

type Statistics struct {
	data   map[string]color.RGBA
	colors [3]color.RGBA
}

func ReadStatistics(r io.Reader) *Statistics {
	stats := new(Statistics)
	stats.data = make(map[string]color.RGBA, 0)

	reader := bufio.NewScanner(r)

	// Skipping the header
	reader.Scan()

	for reader.Scan() {
		row := strings.Split(reader.Text(), ",")

		r1, err := strconv.Atoi(row[1])
		if err != nil {
			return nil
		}

		r2, err := strconv.Atoi(row[2])
		if err != nil {
			return nil
		}

		r3, err := strconv.Atoi(row[3])
		if err != nil {
			return nil
		}

		sum := r1 + r2 + r3

		r := uint8(r1 / sum)
		g := uint8(r2 / sum)
		b := uint8(r3 / sum)

		stats.data[row[0]] = color.RGBA{255 * r, 255 * g, 255 * b, 255}
	}

	return stats
}

func (s *Statistics) GetSubregionColor(subregion string) color.RGBA {
	if v, ok := s.data[subregion]; ok {
		return v
	}
	return color.RGBA{0, 0, 0, 255}
}
