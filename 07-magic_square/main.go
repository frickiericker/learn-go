package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/frickiericker/learn-go/07-magic_square/matrix"
	"github.com/pkg/profile"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	defer profile.Start(profile.CPUProfile).Stop()

	n := 5
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	solver := NewMagicSquareSolver(n, rnd)
	solver.Randomize()

	for {
		if solver.SearchNeighbor() == 0 {
			break
		}
	}
	matrix.Print(solver.Get())

	return nil
}
