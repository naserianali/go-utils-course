package main

import (
	"Curse/utils"
	"fmt"
)

func main() {
	/*checkReverseNumbers()*/
	/*checkRomanToInt()*/
	/*longestCommonPrefix()*/
	/*checkBracketAreValid()*/
	/*removeTargetFromArray()*/
	/*findTheIndexOfTheFirstOccurrenceInAString()*/
	/*searchInsertPosition()*/
	println(utils.LengthOfLastWord("  Go  language  makes  slice per space  "))
}

func checkReverseNumbers() {
	println(utils.CheckReverseNumberIsEqual(121))
	println(utils.CheckReverseNumberIsEqual(-121))
	println(utils.CheckReverseNumberIsEqual(10))
}
func checkRomanToInt() {
	fmt.Printf("%s is equal to %d \n", "III", utils.RomanToInt("III"))
	fmt.Printf("%s is equal to %d \n", "LVIII", utils.RomanToInt("LVIII"))
	fmt.Printf("%s is equal to %d \n", "MCMXCIV", utils.RomanToInt("MCMXCIV"))
}
func longestCommonPrefix() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	println(utils.LongestCommonPrefix(strs1))
	println(utils.LongestCommonPrefix(strs2))
}
func checkBracketAreValid() {
	fmt.Println(utils.CheckValidSigns("()"))
	fmt.Println(utils.CheckValidSigns("()[]{}"))
	fmt.Println(utils.CheckValidSigns("(]"))
	fmt.Println(utils.CheckValidSigns("([])"))
	fmt.Println(utils.CheckValidSigns("([)]"))
}
func removeTargetFromArray() {
	nums := []int{3, 2, 2, 3}
	utils.RemoveTargetFromArray(nums, 2)
}

func findTheIndexOfTheFirstOccurrenceInAString() {
	println(utils.FindTheIndexOfTheFirstOccurrenceInAString("leetcode", "leeto"))
}

func searchInsertPosition() {
	println(utils.SearchInsertPosition([]int{1, 3, 5, 6}, 4))
}
