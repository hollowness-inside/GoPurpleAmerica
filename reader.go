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
	var f zip.File
	found := false
	sname := string(state) + ".txt"
	for _, f = range r.rc.File {
		if f.Name == sname {
			found = true
		}
	}

	if !found {
		return nil, ErrStateName
	}

	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	scanner := Scanner{bufio.NewScanner(rc)}
	reg, err := scanner.ReadRegion()
	return reg, err

}

func (r *ReadCloser) Close() {
	r.rc.Close()
}
