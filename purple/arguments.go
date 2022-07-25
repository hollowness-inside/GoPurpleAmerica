package purple

import (
	"archive/zip"
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

var ErrStateName = errors.New("unknown state name")

type Arguments struct {
	StateName  string
	StatesPath string
	Year       string
	StatsPath  string
	ColorsPath string
	OutputPath string

	Scale       string
	StrokeWidth string
	StrokeColor string
}

func (args *Arguments) Evaluate() (*Purple, error) {
	p := new(Purple)
	p.UseDefaults()

	if args.StatesPath != "" {
		state, err := zipOpen(args.StatesPath, args.StateName, ReadState)
		if err != nil {
			return nil, err
		}

		p.Region = state.(*State)
	}

	p.Year = args.Year

	if args.StatsPath != "" {
		stats, err := zipOpen(args.StatesPath, args.StateName, ReadStatistics)
		if err != nil {
			return nil, err
		}

		p.Stats = stats.(map[string]RGBA)
	}

	if args.StrokeWidth != "" {
		v, err := strconv.ParseFloat(args.StrokeWidth, 64)
		if err != nil {
			return nil, err
		}
		p.StrokeWidth = v
	}

	if args.StrokeColor != "" {
		split := strings.Split(args.StrokeColor, ",")
		rgba, err := AtoiMany(split...)
		if err != nil {
			return nil, err
		}



		p.StrokeColor = RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}

	if args.ColorsPath != "" && args.StatsPath != "" {
		f, err := os.Open(args.ColorsPath)
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

	p.OutputPath = args.OutputPath

	if args.Scale != "" {
		v, err := strconv.ParseFloat(args.Scale, 64)
		if err != nil {
			return nil, err
		}

		p.Scale = v
	}

	return p, nil
}

func zipOpen(filepath, name string, read func(r io.Reader) any) (any, error) {
	reader, err := zip.OpenReader(filepath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var zipFile *zip.File
	for _, v := range reader.File {
		if v.Name == name+".txt" {
			zipFile = v
		}
	}

	f, err := zipFile.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return read(f), nil
}

func AtoiMany(many ...string) ([]int, error) {
	res := make([]int, len(many))
	for i, str := range many {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		res[i] = v
	}

	return res, nil
}

