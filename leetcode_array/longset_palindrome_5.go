package leetcode_array

//https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/

func LongestPalindrome(s string) string {

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := palindrome(s, i, i)
		left2, right2 := palindrome(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start:end]
}

func palindrome(s string, i int, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return i + 1, j
}
