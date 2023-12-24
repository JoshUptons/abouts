package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{1, 2, -3, -4, -5, 6, -7, 2, -3, -4, 1, 2, -5, 3, 2, -1, 5, 3, -2}))
}

func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

func findMaxSubArray(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2
	leftMax := findMaxSubArray(nums, left, mid)
	rightMax := findMaxSubArray(nums, mid+1, right)
	midMax := findCrossing(nums, left, right, mid)
	return max(leftMax, rightMax, midMax)
}

func findCrossing(nums []int, left, right, mid int) int {
	sum := 0
	leftSum := -1 << 31
	rightSum := -1 << 31
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}
	sum = 0
	for i := mid; i < right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		maxVal = v
	}

	return maxVal
}
