package leetcode_array

func LengthOfLongestSubstring(s string) int {

	max := 0
	length := 0

	charMap := make(map[rune]int)
	startIndex := 0

	for k, v := range s {

		if index, ok := charMap[v]; ok == false {
			length++
			charMap[v] = k
		} else {
			//如果存在且在子串范围内
			if startIndex <= index && index <= startIndex+length {

				//结算max
				if max < length {
					max = length
				}

				//这里是关键步骤，找到最新的起始点，以及长度
				startIndex = index + 1
				length = k - index

			} else {
				length++
			}

			charMap[v] = k
		}
	}

	//结算max
	if max < length {
		max = length
	}

	return max
}
