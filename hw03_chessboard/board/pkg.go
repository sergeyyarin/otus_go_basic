package board

import (
	"fmt"
)

func isEven(v int) bool {
	return v%2 == 0
}

func isOdd(v int) bool {
	return !isEven(v)
}

func Build(size int) {
	const (
		uscr = "_"
		hash = "#"
	)

	for i := 0; i < size; i++ {
		var row string
		for j := 0; j < size; j++ {
			if isOdd(i) {
				if isOdd(j) {
					row += uscr
				} else {
					row += hash
				}
			} else {
				if isOdd(j) {
					row += hash
				} else {
					row += uscr
				}
			}
		}
		fmt.Println(row)
	}
}
