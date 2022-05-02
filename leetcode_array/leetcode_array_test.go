package leetcode_array

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{1, 3, 4, 6}
	target := 9

	result := TwoSum(nums, target)

	for _, v := range result {
		if v != 3 && v != 1 {
			t.Error("two sum error")
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	str1 := "babad"
	s := LongestPalindrome(str1)

	if s != "bab" {
		t.Error("get longest substring error")
	}
}
