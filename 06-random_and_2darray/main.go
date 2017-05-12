package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 3
	table := NewMatrix(n, n)
	for i := range table.data {
		table.data[i] = float64(i + 1)
	}
	shuffle(table.data, rnd)

	table.Print()
}

func shuffle(array []float64, rnd *rand.Rand) {
	n := len(array)
	for i := n - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

type Matrix struct {
	data []float64
	rows, cols int
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		data: make([]float64, rows * cols),
		rows: rows,
		cols: cols,
	}
}

func (mat *Matrix) locate(i, j int) int {
	return i * mat.rows + j
}

func (mat *Matrix) Get(i, j int) float64 {
	return mat.data[mat.locate(i, j)]
}

func (mat *Matrix) Set(i, j int, value float64) {
	mat.data[mat.locate(i, j)] = value
}

func (mat *Matrix) Print() {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			if j > 0 {
				fmt.Print("\t")
			}
			fmt.Printf("%.4g", mat.Get(i, j))
		}
		fmt.Print("\n")
	}
}
