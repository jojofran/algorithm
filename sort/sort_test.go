package sort

import "testing"

func TestInsertSort(t *testing.T) {
	nums := []int{15, 0, 2, 4, 7, 9, 1, 0, 0, 1, 0}

	res := InsertSort(nums)

	for k, v := range res {
		if k == 0 {
			if v != 0 {
				t.Error("")
			}
		}
	}
}

func TestSelectSort(t *testing.T) {
	nums := []int{15, 0, 2, 4, 7, 9, 1, 0, 0, 1, 0}
	res := SelectSort(nums)

	for k, v := range res {
		if k == 0 {
			if v != 0 {
				t.Error("")
			}
		}
	}
}
