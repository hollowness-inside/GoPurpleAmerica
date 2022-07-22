package main

import (
	"archive/zip"
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

func (r *ReadCloser) Close() {
	r.rc.Close()
}
