package purple

import (
	"archive/zip"
	"bufio"
	"image/color"
	"os"
	"strconv"
	"strings"
)

type Raw struct {
	County         string
	DataPath       string
	RegionsPath    string
	Year           string
	ColorTablePath string

	StrokeWidth string
	StrokeColor string
}

func (r *Raw) Evaluate() (*Purple, error) {
	p := new(Purple)
	p.UseDefault()

	if r.County != "" {
		p.county = r.County
	}

	if r.DataPath != "" {
		reader, err := zip.OpenReader(r.DataPath)
		if err != nil {
			return nil, err
		}

		p.dataArchive = reader
	}

	if r.Year != "" {
		year, err := strconv.Atoi(r.Year)
		if err != nil {
			return nil, err
		}
		p.year = year
	}

	if r.StrokeWidth != "" {
		v, err := strconv.ParseFloat(r.StrokeWidth, 64)
		if err != nil {
			return nil, err
		}
		p.strokeWidth = v
	}

	if r.StrokeColor != "" {
		split := strings.Split(r.StrokeColor, ",")
		r, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}

		g, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		b, err := strconv.Atoi(split[2])
		if err != nil {
			return nil, err
		}

		a, err := strconv.Atoi(split[3])
		if err != nil {
			return nil, err
		}

		p.strokeColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}

	{
		var reader *zip.ReadCloser
		var err error

		if r.RegionsPath != "" {
			reader, err = zip.OpenReader(r.RegionsPath)
		} else {
			reader, err = zip.OpenReader(r.RegionsPath)
		}

		if err != nil {
			return nil, err
		}
		p.regionsArchive = reader
	}

	if r.ColorTablePath != "" {
		f, err := os.Open(r.ColorTablePath)
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
