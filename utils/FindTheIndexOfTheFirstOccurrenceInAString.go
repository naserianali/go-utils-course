package utils

import "strings"

func FindTheIndexOfTheFirstOccurrenceInAString(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
