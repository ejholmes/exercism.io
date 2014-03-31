package leap

// Returns true of the provided year is a leap year.
func IsLeapYear(year int) bool {
	isDivisibleBy := func(divisor int) bool {
		return year % divisor == 0
	}

	if isDivisibleBy(4) {
		if isDivisibleBy(100) {
			if isDivisibleBy(400) {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}
