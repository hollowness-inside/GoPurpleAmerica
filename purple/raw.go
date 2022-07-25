package purple

import (
	"archive/zip"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Raw struct {
	RegionName  string
	RegionsPath string
	DataPath    string
	Year        string
	ColorsPath  string
	OutputPath  string

	Scale       string
	StrokeWidth string
	StrokeColor string
}

func (r *Raw) UseDefaults() {
	r.RegionName = "USA"
	r.Scale = "10"
	r.StrokeColor = "0,0,0,255"
	r.StrokeWidth = "0.01"
}

func (r *Raw) Evaluate() (*Purple, error) {
	p := new(Purple)
	p.UseDefaults()

	if r.RegionsPath != "" {
		reader, err := zip.OpenReader(r.RegionsPath)
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		var zipFile *zip.File
		for _, f := range reader.File {
			if f.Name == r.RegionName+".txt" {
				zipFile = f
			}
		}

		f, _ := zipFile.Open()
		defer f.Close()

		p.Region = ReadRegion(f)
	}

	p.Year = r.Year

	if r.StrokeWidth != "" {
		v, err := strconv.ParseFloat(r.StrokeWidth, 64)
		if err != nil {
			return nil, err
		}
		p.StrokeWidth = v
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

		p.StrokeColor = RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}

	if r.DataPath != "" {
		reader, err := zip.OpenReader(r.DataPath)
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		var zipFile *zip.File
		for _, f := range reader.File {
			if f.Name == r.RegionName+p.Year+".txt" {
				zipFile = f
			}
		}

		if zipFile == nil {
			return nil, ErrRegionName
		}

		f, _ := zipFile.Open()
		defer f.Close()

		p.Stats = ReadStatistics(f)
	}

	if r.ColorsPath != "" && r.DataPath != "" {
		f, err := os.Open(r.ColorsPath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		colors := [3]RGBA{}
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

			a, err := strconv.Atoi(rgb[3])
			if err != nil {
				return nil, err
			}

			colors[i] = RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		}
	}

	p.OutputPath = r.OutputPath

	if r.Scale != "" {
		v, err := strconv.ParseFloat(r.Scale, 64)
		if err != nil {
			return nil, err
		}

		p.Scale = v
	}

	return p, nil
}
