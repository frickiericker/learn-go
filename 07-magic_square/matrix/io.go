package matrix

import (
	"fmt"
)

func Print(mat Interface) {
	for i := 0; i < mat.Rows(); i++ {
		for j := 0; j < mat.Cols(); j++ {
			if j > 0 {
				fmt.Print("\t")
			}
			fmt.Printf("%.3g", mat.Get(i, j))
		}
		fmt.Print("\n")
	}
}
