package board

import (
	"fmt"
)

const (
	unsc = "_"
	hash = "#"
)

func isEven(v int) bool {
	return v%2 == 0
}

func isOdd(v int) bool {
	return !isEven(v)
}

func getOddValue(index int) string {
	if isOdd(index) {
		return unsc
	}
	return hash
}

func getEvenValue(index int) string {
	if isEven(index) {
		return unsc
	}
	return hash
}

func Build(size int) {
	for i := 0; i < size; i++ {
		var row string
		for j := 0; j < size; j++ {
			if isOdd(i) {
				row += getOddValue(j)
				continue
			}
			row += getEvenValue(j)
		}
		fmt.Println(row)
	}
}
