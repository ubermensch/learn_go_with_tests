package main

func Sum(numbers []int) int {
	sum := 0
	for _, curr := range numbers {
		sum += curr
	}
	return sum
}
