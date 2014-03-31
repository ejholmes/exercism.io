package hamming

func Distance(a, b string) int {
	return countDifferences(a, b, shortest(a, b))
}

func countDifferences(a, b string, length int) int {
	var differences = 0

	for idx := 0; idx < length; idx++ {
		if a[idx] != b[idx] {
			differences++
		}
	}

	return differences
}

func shortest(a, b string) int {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	return length
}
