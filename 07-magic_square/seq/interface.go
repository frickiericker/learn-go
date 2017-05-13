package seq

type Interface interface {
	Len() int
}

type Swappable interface {
	Interface
	Swap(i, j int)
}
