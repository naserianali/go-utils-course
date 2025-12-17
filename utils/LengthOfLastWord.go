package utils

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	myRune := strings.Fields(s)
	return len(myRune[len(myRune)-1])
}
