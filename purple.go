package main

type Purple struct {
	county string
	colors []Color
	data   string
	year   int
}

func NewPurple() Purple {
	p := Purple{}
	p.colors = []Color{
		Color{255, 0, 0},
		Color{0, 255, 0},
		Color{0, 0, 255},
	}

	return p
}
