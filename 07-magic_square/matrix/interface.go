package matrix

type Interface interface {
	Rows() int
	Cols() int
	Get(i, j int) float64
	Set(i, j int, value float64)
}
