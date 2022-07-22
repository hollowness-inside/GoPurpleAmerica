package purple

import (
	"archive/zip"
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
	defer f.Close()

	s := NewScanner(f)
	counties <- s.ScanCounty()
}
