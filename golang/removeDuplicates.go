package main

import "fmt"

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// expected output []nums{0, 1, 2, 3, 4} which would be 5 elements
	// the return value of the function is the number of elements
	output := removeDuplicates(input)
	// expected 5

	fmt.Println(output)
	input2 := []int{1, 1, 2}
	output = removeDuplicates(input2)
	fmt.Println(output)
}

func removeDuplicates(nums []int) int {

	left := 1
	right := 1

	for right < len(nums) {
		if nums[right] == nums[right-1] {
		} else {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	fmt.Println(nums)
	return left
}
