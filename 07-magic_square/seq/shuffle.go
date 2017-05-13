package seq

import (
	"math/rand"
)

func Shuffle(seq Swappable, rnd *rand.Rand) {
	for i := seq.Len() - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		seq.Swap(i, j)
	}
}
