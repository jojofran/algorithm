package main

import (
	"github/jojofran/algorithm/leetcode_array"
)

func main() {
	str := "dvdf"
	_ = leetcode_array.LengthOfLongestSubstring(str)

	str1 := "babad"
	s := leetcode_array.LongestPalindrome(str1)
	print(s)
}
