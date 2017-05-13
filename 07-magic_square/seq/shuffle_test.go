package seq

import (
	"math/rand"
	"testing"
)

func TestShuffle_isPermutation(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))
	seq := Float64Slice{[]float64{1, 2, 3, 4}}
	Shuffle(seq, rnd)

	counter := map[float64]int{}
	for _, value := range seq.slice {
		counter[value] += 1
	}
	for _, count := range counter {
		if count != 1 {
			t.Errorf("%v", seq)
		}
	}
}

func TestShuffle_allowEmptySeq(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))
	seq := Float64Slice{[]float64{}}
	Shuffle(seq, rnd)
}
