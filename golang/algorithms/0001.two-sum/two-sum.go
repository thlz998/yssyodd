package problem0001

// 思路一
func twoSum(nums []int, target int) []int {
	tnums := nums
	for i, b := range nums {
		tnums = tnums[1:]
		for j, c := range tnums {
			if b+c == target {
				println(i, j+i+1)
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

// 思路二
func twoSum2(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			println(j, i)
			return []int{j, i}
			// 注意，顺序是j，i
		}
		// 把b和i的值，存入map
		index[b] = i
	}
	return nil
}
