package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MrPythoneer/nifty/purple_america/purple"
)

func Help() {
	fmt.Println(`
	purple [options] -r regions.zip -o output.svg
	Options:
	\t-c/--county COUNTY_ABBR\tSelects county to draw (USA by default)
	\t-d/--data path.zip\tSelects archive containing statisics data
	\t-r/--region path.zip\tSelects archive containing region points data
	\t-y/--year INT\tSuffix of the state statistics file name
	\t-n/--colors filepath\tDraws with colors presented in the file
	\t-N/--new-color-table output\tSaves an example of a color file
	\t--sw/--stroke-width FLOAT\tSets stroke width
	\t--sc/--stroke-color R,G,B,A\tSets stroke color
	\t-s/--scale FLOAT\tScales the result by the given factor (10 by default)`)
}

func main() {
	// TODO: Map scale
	// TODO: County abbreviation or full name
	// TODO: Geometric data by county
	// TODO: Geometric data by several counties
	// TODO: Election results by county
	// TODO: Election results by several counties
	// TODO: Colors lookup table
	// TODO: Use different data sets

	// Draws outline of the county
	// purple --county COUNTY
	// purple -c COUNTY

	// Draws purple map representing election results
	// purple --county COUNTY --data elections.zip --year 2015
	// purple -c COUNTY -d elections.zip -y 2015

	// Creates colors.txt, lookup color table
	// purple --new-color-table [colors.txt]
	// purple -N [colors.txt]

	// Uses the specified lookup color table
	// purple --colors colors.txt
	// purple -n colors.txt

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
		Help()
		return
	}

	p, err := pRaw.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	p.Draw()
}

func createColorTable(output string) {

}
