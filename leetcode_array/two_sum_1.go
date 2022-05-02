package leetcode_array

//https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	resultDic := make(map[int]int)

	for k, v := range nums {

		if val, ok := resultDic[target-v]; ok {
			return []int{k, val}
		}
		resultDic[v] = k
	}

	return nil
}
