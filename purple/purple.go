package purple

import "archive/zip"

type Purple struct {
	counties []string
	year     int
	colors   [3][3]int

	dataArchive    *zip.ReadCloser
	regionsArchive *zip.ReadCloser
}

func (p *Purple) GetCounties() ([]*County, error) {
	ch := make(chan *County)
	err := p.getCounties(ch)
	if err != nil {
		return err
	}

	counties := make([]*County, 0)
	for i, c := range ch {
		counties[i] = c
	}

	return counties
}

func (p *Purple) getCounties(counties chan *County) error {
	files := make(map[string]*zip.File, 0)
	for _, county := range p.counties {
		files[county] = nil
	}

	for _, v := range p.regionsArchive.File {
		if _, ok := files[v.Name]; ok {
			files[v.Name] = v
		}
	}

	if len(files) != len(p.counties) {
		return nil, ErrCountyName
	}

	for _, file := range files {
		go func() {
			counties <- readFile(file)
		}()
	}

	close(counties)

	return nil
}
