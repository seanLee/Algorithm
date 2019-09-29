package main

import "math"

func lengthOfLongestSubstring(s string) int {
	count := len(s)

	begin := 0
	end := 0

	res := ""
	storage := map[string]int{}

	for index := 1; index < count; index++ {
		key := string(s[index])
		if _, ok := storage[key]; ok {
			last := storage[key]
			begin = int(math.Max(float64(last), float64(begin)))
		}
		end += 1
		storage[key] = end

		current := s[begin:end]
		if len(current) >= len(res) {
			res = current
		}
	}

	return len(res)
}