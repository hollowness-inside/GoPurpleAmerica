package main

import (
	"archive/zip"
)

type Reader struct {
	r *zip.ReadCloser
}

func OpenReader(name string) (*Reader, error) {
	rc, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	return &Reader{rc}, nil
}
