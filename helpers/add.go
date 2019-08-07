package helpers

import "fmt"

func Add(nums ...int) int {
	sum := 0

	fmt.Println("hello")

	for _, n := range nums {
		sum += n
	}

	return sum
}

func Subtract(nums ...int) int {
	diff := 0

	return diff
}

func Divide(nums ...int) int {
	val := 0

	for _, n := range nums {
		val /= n
	}

	return val
}
