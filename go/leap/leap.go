package leap

// Returns true of the provided year is a leap year.
func IsLeapYear(year int) bool {
	isDivisibleBy := func(divisor int) bool {
		return year % divisor == 0
	}

	return isDivisibleBy(4) && !isDivisibleBy(100) || isDivisibleBy(400)
}
