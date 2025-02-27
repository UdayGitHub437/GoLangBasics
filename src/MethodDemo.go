package main

type square struct {
	side float64
}

func (s square) perim() float64 {
	return 4 * s.side
}

func (s *square) area() float64 {
	return s.side * s.side

}
