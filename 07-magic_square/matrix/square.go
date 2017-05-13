package matrix

type Square struct {
	data []float64
	size int
}

func NewSquare(size int) *Square {
	return &Square{
		data: make([]float64, size*size),
		size: size,
	}
}

func (mat *Square) Data() []float64 {
	return mat.data
}

func (mat *Square) Size() int {
	return mat.size
}

func (mat *Square) Rows() int {
	return mat.size
}

func (mat *Square) Cols() int {
	return mat.size
}

func (mat *Square) locate(i, j int) int {
	return i*mat.size + j
}

func (mat *Square) Get(i, j int) float64 {
	return mat.data[mat.locate(i, j)]
}

func (mat *Square) Set(i, j int, value float64) {
	mat.data[mat.locate(i, j)] = value
}
