package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MrPythoneer/nifty/purple_america/purple"
)

const HelpMsg = `purple [options] -r regions.zip -o output.svg
Options:
	-c/--county COUNTY_ABBR	Selects county to draw (USA by default)
	-d/--data path.zip	Selects archive containing statisics data
	-r/--region path.zip	Selects archive containing region points data
	-y/--year INT	Suffix of the state statistics file name
	-n/--colors filepath	Draws with colors presented in the file
	-N/--new-color-table output	Saves an example of a color file
	-s/--scale FLOAT	Scales the result by the given factor (10 by default)
	--sw/--stroke-width FLOAT	Sets stroke width
	--sc/--stroke-color R,G,B,A	Sets stroke color`

func main() {
	// TODO: County abbreviation or full name
	// TODO: Election results by county
	// TODO: Election results by several counties
	// TODO: Use different data sets

	pRaw := purple.Raw{}

	i := 1

	for i < len(os.Args) {
		switch os.Args[i] {
		case "-c", "--county":
			i++
			pRaw.County = os.Args[i]
		case "-d", "--data":
			i++
			pRaw.DataPath = os.Args[i]
		case "-r", "--region":
			i++
			pRaw.RegionsPath = os.Args[i]
		case "-o", "--output":
			i++
			pRaw.OutputPath = os.Args[i]
		case "-y", "--year":
			i++
			pRaw.Year = os.Args[i]
		case "-n", "--colors":
			i++
			pRaw.ColorTablePath = os.Args[i]
		case "-N", "--new-color-table":
			i++
			createColorTable(os.Args[i])
			return
		case "--sw", "--stroke-width":
			i++
			pRaw.StrokeWidth = os.Args[i]
		case "--sc", "--stroke-color":
			i++
			pRaw.StrokeColor = os.Args[i]
		case "-s", "--scale":
			i++
			pRaw.Scale = os.Args[i]
		}

		i++
	}

	if pRaw.RegionsPath == "" || pRaw.OutputPath == "" {
		fmt.Println("Please provide output file name and regions data archive")
		fmt.Println(HelpMsg)
		return
	}

	p, err := pRaw.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	p.Draw()
}

func createColorTable(output string) {
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString("255 0 0\n0 255 0\n0 0 255")
	if err != nil {
		log.Fatal(err)
	}
}
