package anagram

import(
	"strings"
	"sort"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	anagrams := make([]string, 0)

	for idx, candidate := range candidates {
		candidates[idx] = strings.ToLower(candidate)
	}

	for _, candidate := range candidates {
		if subject != candidate && sorted(subject) == sorted(candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func sorted(str string) string {
	characters := strings.Split(str, "")
	sort.Strings(characters)
	return strings.Join(characters, "")
}
