package main

import (
	"fmt"
	"log"
	"os"
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

	pRaw := PurpleRaw{}

	i := 1

	for i < len(os.Args) {
		switch os.Args[i] {
		case "-c", "--county":
			i++
			pRaw.counties = os.Args[i]
		case "-d", "--data":
			i++
			pRaw.dataPath = os.Args[i]
		case "-y", "--year":
			i++
			pRaw.year = os.Args[i]
		case "-n", "--colors":
			i++
			pRaw.colorTablePath = os.Args[i]
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

	fmt.Println(*p)
}

func createColorTable(a string) {

}
