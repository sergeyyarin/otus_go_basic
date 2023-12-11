package board

import (
	"errors"
)

const (
	unsc = "_"
	hash = "#"
)

var ErrInvalidValue = errors.New("invalid value")

func isOdd(v int) bool {
	return v%2 != 0
}

func getOddValue(index int) string {
	if isOdd(index) {
		return unsc
	}
	return hash
}

func getEvenValue(index int) string {
	if !isOdd(index) {
		return unsc
	}
	return hash
}

func Build(size int) (string, error) {
	var res string

	if size <= 0 || isOdd(size) {
		return res, ErrInvalidValue
	}

	for i := 0; i < size; i++ {
		var row string
		for j := 0; j < size; j++ {
			if isOdd(i) {
				row += getOddValue(j)
				continue
			}
			row += getEvenValue(j)
		}
		res += row
		res += "\n"
	}

	return res, nil
}
