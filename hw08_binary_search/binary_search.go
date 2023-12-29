package binarysearch

func BinarySearch(s []int, v int) bool {
	if len(s) == 0 {
		return false
	}

	b, e := 0, len(s)

	for b <= e {
		m := (b + e) / 2

		if s[m] < v {
			b = m + 1
		} else {
			e = m - 1
		}
	}

	if b == len(s) || s[b] != v {
		return false
	}

	return true
}
