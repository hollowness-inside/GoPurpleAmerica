package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type PurpleRaw struct {
	counties       string
	dataPath       string
	year           string
	colorTablePath string
}

func (p *PurpleRaw) Evaluate() (*Purple, error) {
	purple := new(Purple)
	purple.counties = []string{"USA"}
	purple.colors = [3][3]int{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
	}

	if p.counties != "" {
		counties := strings.Split(p.counties, ",")
		purple.counties = counties
	}

	purple.dataPath = p.dataPath

	if p.year != "" {
		year, err := strconv.Atoi(p.year)
		if err != nil {
			return nil, err
		}
		purple.year = year
	}

	if p.colorTablePath != "" {
		f, err := os.Open(p.colorTablePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		colors := [3][3]int{}
		for i := 0; i < 3; i++ {
			if !(sc.Scan() && sc.Err() == nil) {
				return nil, sc.Err()
			}

			rgbText := sc.Text()
			rgb := strings.Split(rgbText, " ")
			r, err := strconv.Atoi(rgb[0])
			if err != nil {
				return nil, err
			}

			g, err := strconv.Atoi(rgb[1])
			if err != nil {
				return nil, err
			}

			b, err := strconv.Atoi(rgb[2])
			if err != nil {
				return nil, err
			}

			colors[i] = [3]int{r, g, b}
		}

		purple.colors = colors
	}

	return purple, nil
}

type Purple struct {
	counties []string
	dataPath string
	year     int
	colors   [3][3]int
}
