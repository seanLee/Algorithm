package main

func bubbleSort(numbers []int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			left := numbers[j]
			right := numbers[j+1]
			if left > right {
				temp := left
				numbers[j] = right
				numbers[j+1] = temp
			}
		}
	}

	return numbers
}