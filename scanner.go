package main

import (
	"bufio"
	"strconv"
)

type Scanner struct {
	sc *bufio.Scanner
}

func (s *Scanner) ReadFloat64() (float64, error) {
	s.sc.Scan()
	return strconv.ParseFloat(s.sc.Text(), 64)
}

func (s *Scanner) ReadInt() (int, error) {
	s.sc.Scan()
	return strconv.Atoi(s.sc.Text())
}

func (s *Scanner) ReadPoint() (Point, error) {
	x, err := s.ReadFloat64()
	if err != nil {
		return Point{}, err
	}

	y, err := s.ReadFloat64()
	if err != nil {
		return Point{}, err
	}

	return Point{x, y}, nil
}

func (s *Scanner) ReadBBox() (BBox, error) {
	p1, err := s.ReadPoint()
	if err != nil {
		return BBox{}, err
	}

	p2, err := s.ReadPoint()
	if err != nil {
		return BBox{}, err
	}

	return BBox{p1, p2}, nil
}

func (s *Scanner) ReadRegion() (*Region, error) {
	bbox, err := s.ReadBBox()
	if err != nil {
		return nil, err
	}

	n, err := s.ReadInt()
	if err != nil {
		return nil, err
	}

	subregs := make([]*Subregion, n)
	for i := 0; i < n; i++ {
		subreg, err := s.ReadSubregion()
		if err != nil {
			return nil, err
		}

		subregs[i] = subreg
	}

	return &Region{bbox, n, subregs}, nil
}

func (s *Scanner) ReadSubregion() (*Subregion, error) {
	return nil, nil
}
