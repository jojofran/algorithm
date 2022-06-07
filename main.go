package main

import (
	"fmt"
	"github/jojofran/algorithm/leetcode_array"
	"github/jojofran/algorithm/sort"
)

func main() {
	str := "dvdf"
	_ = leetcode_array.LengthOfLongestSubstring(str)

	str1 := "babad"
	s := leetcode_array.LongestPalindrome(str1)
	print(s)

	nums := []int{9, 1, 3, 4, 7, 2, 1, 0, 5, 10, 13, 0, 0}

	ret := sort.InsertSort(nums)

	for k, v := range ret {
		fmt.Printf("index: %d, value: %d\n", k, v)
	}
}
