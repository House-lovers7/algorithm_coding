package main

import (
	"fmt"
	"math/rand"
)

func partition(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func quickSort(numbers []int) []int {
	var sort func([]int, int, int)
	sort = func(numbers []int, low int, high int) {
		if low < high {
			partitionIndex := partition(numbers, low, high)
			sort(numbers, low, partitionIndex-1)
			sort(numbers, partitionIndex+1, high)
		}
	}

	sort(numbers, 0, len(numbers)-1)
	return numbers
}

func main() {
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(1000))
	}
	fmt.Println(quickSort(nums))
}
