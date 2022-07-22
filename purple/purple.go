package purple

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Raw struct {
	Counties       string
	DataPath       string
	Year           string
	ColorTablePath string
}

func (pr *Raw) Evaluate() (*Purple, error) {
	p := new(Purple)
	p.counties = []string{"USA"}
	p.colors = [3][3]int{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
	}

	if pr.Counties != "" {
		counties := strings.Split(pr.Counties, ",")
		p.counties = counties
	}

	p.dataPath = pr.DataPath

	if pr.Year != "" {
		year, err := strconv.Atoi(pr.Year)
		if err != nil {
			return nil, err
		}
		p.year = year
	}

	if pr.ColorTablePath != "" {
		f, err := os.Open(pr.ColorTablePath)
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

		p.colors = colors
	}

	return p, nil
}

type Purple struct {
	counties []string
	dataPath string
	year     int
	colors   [3][3]int
}
