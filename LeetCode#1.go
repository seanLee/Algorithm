package main

func twoSum(nums []int, target int) []int {
	maps := map[int]int{}

	for i:=0 ;i < len(nums) ;i++  {
		number := nums[i]
		if first, key := maps[number]; key {
			return []int{first, i}
		} else {
			maps[target-number] = i
		}
	}

	return []int{}
}