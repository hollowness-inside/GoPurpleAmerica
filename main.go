package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MrPythoneer/nifty/purple_america/purple"
)

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
			pRaw.Counties = os.Args[i]
		case "-d", "--data":
			i++
			pRaw.DataPath = os.Args[i]
		case "-y", "--year":
			i++
			pRaw.Year = os.Args[i]
		case "-n", "--colors":
			i++
			pRaw.ColorTablePath = os.Args[i]
		case "--nct", "--new-color-table":
			createColorTable(os.Args[i])
			return
		}

		i++
	}

	p, err := pRaw.Evaluate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}

func createColorTable(a string) {

}
