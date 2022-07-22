package main

import (
	"archive/zip"
	"bufio"
)

type ReadCloser struct {
	rc *zip.ReadCloser
}

func OpenReader(name string) (*ReadCloser, error) {
	rc, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	return &ReadCloser{rc}, nil
}

func (r *ReadCloser) GetRegion(state StateName) (*Region, error) {
	sname := string(state) + ".txt"
	for _, f := range r.rc.File {
		if f.Name == sname {
			return readRegion(f)
		}
	}

	return nil, ErrStateName
}

func readRegion(f zip.File) (*Region, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	scanner := Scanner{bufio.NewScanner(rc)}
	bbox, err := scanner.ReadBBox()
	if err != nil {
		return nil, err
	}

	n, err := scanner.ReadInt()
	if err != nil {
		return nil, err
	}

	subregs := make([]Subregion, n)
	for i := 0; i < n; i++ {
		subregs[i] = readSubregion()
	}

	return &Region{bbox, n, subregs}, nil
}

func (r *ReadCloser) Close() {
	r.rc.Close()
}
