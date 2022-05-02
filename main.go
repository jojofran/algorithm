package main

import "github/jojofran/algorithm/array"

func main() {
	str := "dvdf"
	_ = array.LengthOfLongestSubstring(str)

	str1 := "babad"
	s := array.LongestPalindrome(str1)
	print(s)
}
