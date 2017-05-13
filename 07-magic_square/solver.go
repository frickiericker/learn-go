package main

import (
	"math"
	"math/rand"

	"github.com/frickiericker/learn-go/07-magic_square/matrix"
	"github.com/frickiericker/learn-go/07-magic_square/seq"
)

type MagicSquareSolver struct {
	square *matrix.Square
	magic  float64
	rnd    *rand.Rand
}

func NewMagicSquareSolver(size int, rnd *rand.Rand) *MagicSquareSolver {
	square := matrix.NewSquare(size)
	for i := range square.Data() {
		square.Data()[i] = float64(i + 1)
	}
	return &MagicSquareSolver{
		square: square,
		magic:  float64(magicConstant(size)),
		rnd:    rnd,
	}
}

func (solver *MagicSquareSolver) Get() *matrix.Square {
	return solver.square
}

func (solver *MagicSquareSolver) Randomize() {
	seq.Shuffle(seq.NewFloat64Slice(solver.square.Data()), solver.rnd)
}

func (solver *MagicSquareSolver) evaluate() float64 {
	loss := float64(0)
	magic := solver.magic
	for i := 0; i < solver.square.Size(); i++ {
		loss += math.Abs(sumRow(solver.square, i) - magic)
		loss += math.Abs(sumCol(solver.square, i) - magic)
	}
	loss += math.Abs(sumDiag(solver.square) - magic)
	loss += math.Abs(sumAntidiag(solver.square) - magic)
	return loss
}

func (solver *MagicSquareSolver) SearchNeighbor() float64 {
	n := solver.square.Size()
	data := solver.square.Data()
	c1 := solver.rnd.Intn(n * n)
	c2 := solver.rnd.Intn(n * n)

	lossBefore := solver.evaluate()
	data[c1], data[c2] = data[c2], data[c1]
	lossAfter := solver.evaluate()
	if lossAfter - 3 > lossBefore {
		data[c1], data[c2] = data[c2], data[c1]
		return lossBefore
	}
	return lossAfter
}

func sumRow(square *matrix.Square, row int) float64 {
	sum := float64(0)
	for col := 0; col < square.Cols(); col++ {
		sum += square.Get(row, col)
	}
	return sum
}

func sumCol(square *matrix.Square, col int) float64 {
	sum := float64(0)
	for row := 0; row < square.Rows(); row++ {
		sum += square.Get(row, col)
	}
	return sum
}

func sumDiag(square *matrix.Square) float64 {
	sum := float64(0)
	for i := 0; i < square.Size(); i++ {
		sum += square.Get(i, i)
	}
	return sum
}

func sumAntidiag(square *matrix.Square) float64 {
	sum := float64(0)
	size := square.Size()
	for i := 0; i < size; i++ {
		sum += square.Get(i, size-i-1)
	}
	return sum
}

func magicConstant(n int) int {
	return n * (n*n + 1) / 2
}
