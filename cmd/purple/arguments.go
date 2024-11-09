package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/MrPythoneer/nifty/purple/purple"
)

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

func (args *Arguments) Evaluate() (*purple.Purple, error) {
	p := new(purple.Purple)
	p.UseDefaults()

	if args.StatesPath != "" {
		state, err := zipRead(args.StatesPath, args.StateName, purple.ReadState)
		if err != nil {
			return nil, err
		}

		p.State = state.(*purple.State)
	}

	p.Year = args.Year

	if args.StatsPath != "" {
		stats, err := zipRead(args.StatsPath, args.StateName+args.Year, ReadStatistics)
		if err != nil {
			return nil, err
		}

		p.Stats = stats.(map[string]purple.RGBA)
	}

	if args.ColorsPath != "" && args.StatsPath != "" {
		f, err := os.Open(args.ColorsPath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		colors := [3]purple.RGBA{}
		for i := 0; i < 3; i++ {
			if !(sc.Scan() && sc.Err() == nil) {
				return nil, sc.Err()
			}

			clr, err := ParseRGBA(sc.Text())
			if err != nil {
				return nil, err
			}

			colors[i] = clr
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

	if args.StrokeWidth != "" {
		v, err := strconv.ParseFloat(args.StrokeWidth, 64)
		if err != nil {
			return nil, err
		}
		p.StrokeWidth = v
	}

	if args.StrokeColor != "" {
		clr, err := ParseRGBA(args.StrokeColor)
		if err != nil {
			return nil, err
		}

		p.StrokeColor = clr
	}

	return p, nil
}

func ReadStatistics(r io.Reader) any {
	data := make(map[string]purple.RGBA, 0)

	reader := bufio.NewScanner(r)

	// Skipping the header
	reader.Scan()

	for reader.Scan() {
		row := strings.Split(reader.Text(), ",")

		r1, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil
		}

		r2, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil
		}

		r3, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil
		}

		sum := r1 + r2 + r3

		r := uint8((r1 / sum) * 255)
		g := uint8((r2 / sum) * 255)
		b := uint8((r3 / sum) * 255)

		data[row[0]] = purple.RGBA{r, g, b, 255}
	}

	return data
}

func zipRead(filepath, name string, read func(r io.Reader) any) (any, error) {
	name = name + ".txt"
	reader, err := zip.OpenReader(filepath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var zipFile *zip.File
	for _, v := range reader.File {
		if v.Name == name {
			zipFile = v
			break
		}
	}

	if zipFile == nil {
		return nil, fmt.Errorf("cannot find '%s' in '%s'", name, filepath)
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

func ParseRGBA(text string) (purple.RGBA, error) {
	split := strings.Split(text, ",")
	rgba, err := AtoiMany(split...)
	if err != nil {
		return purple.RGBA{}, err
	}

	r := uint8(rgba[0])
	g := uint8(rgba[1])
	b := uint8(rgba[2])
	a := uint8(rgba[3])

	return purple.RGBA{r, g, b, a}, nil
}
