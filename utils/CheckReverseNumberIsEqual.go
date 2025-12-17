package utils

import (
	"strconv"
)

func CheckReverseNumberIsEqual(x int) bool {
	convertToString := strconv.Itoa(x)
	runeSlice := make([]rune, len(convertToString))
	n := 0
	for _, r := range convertToString {
		runeSlice[n] = r
		n++
	}
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	result := string(runeSlice)
	return result == convertToString
}
