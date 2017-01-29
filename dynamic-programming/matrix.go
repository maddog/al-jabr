package main

type Matrix struct {
	height int
	width  int
	table  []float64
}

func NewMatrix(height, width int) Matrix {
	return Matrix{height, width, make([]float64, height*width)}
}

func (s *Matrix) Set(i, j int, val float64) {
	s.table[s.width*i+j] = val
}

func (s *Matrix) Get(i, j int) float64 {
	return s.table[s.width*i+j]
}
