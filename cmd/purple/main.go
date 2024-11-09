package main

import (
	"fmt"
	"log"
	"os"

	"github.com/llgcode/draw2d/draw2dsvg"
)

// TODO: County abbreviation or full name
// TODO: Election results by county
// TODO: Election results by several counties
// TODO: Use different data sets

const HelpMsg = `Usage: purple [options] -o output.svg
Options:
	-r/--region REGION_NAME		Selects region to draw (USA by default)
	-rd/--regions-data regions.zip	Selects archive containing region points data
	-d/--data path.zip		Selects archive containing statisics data
	-y/--year INT			Suffix of the state statistics file name
	-n/--colors filepath		Draws with colors presented in the file
	-N output			Saves an example of a color file
	-h/--help			Shows this message
	-o/--output image.svg		Output path
	-s/--scale FLOAT		Scales the result by the given factor (10 by default)
	-sw/--stroke-width FLOAT	Sets stroke width
	-sc/--stroke-color R,G,B,A	Sets stroke color`

func main() {
	args := DefaultArguments()

	if len(os.Args) <= 1 {
		fmt.Println(HelpMsg)
		return
	}

	i := 1
	for i < len(os.Args) {
		switch os.Args[i] {
		case "-r", "--region":
			i++
			args.StateName = os.Args[i]
		case "-rd", "--regions-data":
			i++
			args.StatesPath = os.Args[i]
		case "-o", "--output":
			i++
			args.OutputPath = os.Args[i]
		case "-d", "--data":
			i++
			args.StatsPath = os.Args[i]
		case "-y", "--year":
			i++
			args.Year = os.Args[i]
		case "-n", "--colors":
			i++
			args.ColorsPath = os.Args[i]
		case "-N":
			i++
			createColorTable(os.Args[i])
			return
		case "-sw", "--stroke-width":
			i++
			args.StrokeWidth = os.Args[i]
		case "-sc", "--stroke-color":
			i++
			args.StrokeColor = os.Args[i]
		case "-s", "--scale":
			i++
			args.Scale = os.Args[i]
		}
		i++
	}

	if args.StatesPath == "" {
		fmt.Println("Please provide regions data archive")
		return
	}

	if args.OutputPath == "" {
		fmt.Println("Please provide output file path")
		return
	}

	p, err := args.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	svg := p.GenerateSVG()
	draw2dsvg.SaveToSvgFile(args.OutputPath, svg)
}

func createColorTable(output string) {
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString("255 0 0\n0 255 0\n0 0 255"); err != nil {
		log.Fatal(err)
	}
}
