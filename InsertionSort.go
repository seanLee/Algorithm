package main

func insertionSort(numbers []int) []int {

	for i := 1; i < len(numbers); i++ {
		left := numbers[i]
		start := i - 1
		for ; start >= 0; start-- {
			if numbers[start] > left {
				numbers[start+1] = numbers[start]
			} else {
				break
			}
		}
		numbers[start+1] = left
	}

	return numbers
}