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
	Region         string
	RegionsPath    string
	DataPath       string
	Year           string
	ColorTablePath string
	OutputPath     string

	StrokeWidth string
	StrokeColor string
	Scale       string
}

func (r *Raw) Evaluate() (*Purple, error) {
	p := new(Purple)
	p.UseDefault()

	if r.Region != "" {
		p.regionName = r.Region
	}

	if r.DataPath != "" {
		reader, err := zip.OpenReader(r.DataPath)
		if err != nil {
			return nil, err
		}

		p.dataArchive = reader
	}

	if r.RegionsPath != "" {
		reader, err := zip.OpenReader(r.RegionsPath)
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		var zipFile *zip.File
		for _, f := range reader.File {
			if f.Name == p.regionName+".txt" {
				zipFile = f
			}
		}

		f, _ := zipFile.Open()
		defer f.Close()

		p.region = ReadRegion(f)
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

	if r.ColorTablePath != "" {
		f, err := os.Open(r.ColorTablePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		colors := [3]color.RGBA{}
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

			colors[i] = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		}

		p.colors = colors
	}

	p.outputPath = r.OutputPath

	if r.Scale != "" {
		v, err := strconv.ParseFloat(r.Scale, 64)
		if err != nil {
			return nil, err
		}

		p.scale = v
	}

	return p, nil
}
