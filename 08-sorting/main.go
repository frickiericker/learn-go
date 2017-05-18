package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	benchmark(10, 3)
	benchmark(100, 3)
	benchmark(1000, 3)
	benchmark(10000, 3)
}

func benchmark(size int, measurements int) {
	fmt.Printf("size = %d\n", size)
	data := make([]int, size)
	for i := 0; i < measurements; i++ {
		randomize(data)
		start := time.Now()
		sort(data)
		elapsed := time.Since(start)
		if isSorted(data) {
			fmt.Printf("  sorted in %s\n", elapsed)
		} else {
			fmt.Println("  FAILED")
		}
	}
}

func randomize(data []int) {
	for i := range data {
		data[i] = rand.Intn(len(data))
	}
}

func isSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i - 1] {
			return false
		}
	}
	return true
}

func sort(data []int) {
	quickSort(data)
}

func bubbleSort(data []int) {
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		j := i
		for j > 0 && data[j - 1] > tmp {
			data[j] = data[j - 1]
			j--
		}
		data[j] = tmp
	}
}

func quickSort(data []int) {
	if len(data) > 1 {
		p, q := partition(data)
		sort(data[:p])
		sort(data[q:])
	}
}

func partition(data []int) (int, int) {
	n := len(data)
	i := 0
	j := n - 1
	pivot := middle(data[0], data[n/2], data[n-1])
	for {
		for data[i] < pivot {
			i++
		}
		for data[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return i, j + 1
}

func middle(x, y, z int) int {
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	return y
}

func mergeSort(data []int) {
	copy(data, mergeSorted(data))
}

func mergeSorted(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	p := len(data) / 2
	return merge(mergeSorted(data[:p]), mergeSorted(data[p:]))
}

func merge(a, b []int) []int {
	result := make([]int, len(a) + len(b))
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[i + j] = a[i]
			i++
		} else {
			result[i + j] = b[j]
			j++
		}
	}
	for i < len(a) {
		result[i + j] = a[i]
		i++
	}
	for j < len(b) {
		result[i + j] = b[j]
		j++
	}
	return result
}
