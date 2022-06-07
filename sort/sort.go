package sort

// InsertSort
// 插入排序
func InsertSort(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums
}

// SelectSort
// 选择排序
func SelectSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		minindex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minindex] {
				minindex = j
			}
		}

		if i != minindex {
			nums[minindex], nums[i] = nums[i], nums[minindex]
		}
	}

	return nums
}
