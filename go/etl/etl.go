package etl

import "strings"

type Crufty map[int][]string
type Shiny map[string]int

func Transform(input Crufty) Shiny {
	output := make(Shiny)

	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}

	return output
}
