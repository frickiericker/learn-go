package seq

type Float64Slice struct {
	slice []float64
}

func NewFloat64Slice(slice []float64) Float64Slice {
	return Float64Slice{slice}
}

func (seq Float64Slice) Len() int {
	return len(seq.slice)
}

func (seq Float64Slice) Swap(i, j int) {
	seq.slice[i], seq.slice[j] = seq.slice[j], seq.slice[i]
}
