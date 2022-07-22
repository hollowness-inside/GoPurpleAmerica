package purple

import (
	"archive/zip"
	"bufio"
	"strconv"
	"strings"
	"sync"
)

type Purple struct {
	counties []string
	year     int
	colors   [3][3]int

	dataArchive    *zip.ReadCloser
	regionsArchive *zip.ReadCloser
}

func (p *Purple) GetCounties() ([]*County, error) {
	files := make(map[string]*zip.File, 0)
	for _, county := range p.counties {
		files[county+".txt"] = nil
	}

	for _, v := range p.regionsArchive.File {
		if _, ok := files[v.Name]; ok {
			files[v.Name] = v
		}
	}

	if len(files) != len(p.counties) {
		return nil, ErrCountyName
	}

	wg := &sync.WaitGroup{}
	ch := make(chan *County, len(p.counties))

	for _, file := range files {
		wg.Add(1)
		go fileReader(wg, ch, file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	counties := make([]*County, len(p.counties))

	i := 0
	for c := range ch {
		counties[i] = c
		i++
	}

	return counties, nil
}

func fileReader(wg *sync.WaitGroup, counties chan *County, zipFile *zip.File) {
	defer wg.Done()

	f, _ := zipFile.Open()
	sc := bufio.NewScanner(f)
	p := readPoint(sc)

	counties <- &County{p}
}

func readPoint(sc *bufio.Scanner) Point {
	sc.Scan()
	xy := strings.Split(sc.Text(), "   ")
	x, _ := strconv.ParseFloat(xy[0], 64)
	y, _ := strconv.ParseFloat(xy[1], 64)

	return Point{x, y}
}

type Point struct {
	x, y float64
}
