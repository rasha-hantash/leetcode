package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	return helper(nums, 0, len(nums)-1)

}

func helper(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	p := (left + right) / 2
	leftSum := helper(nums, left, p)
	rightSum := helper(nums, p+1, right)
	sumCross := crossSum(nums, left, right, p)
	return MaxOf(leftSum, rightSum, sumCross)
}

func crossSum(nums []int, left int, right int, p int) int {
	if left == right {
		return nums[left]
	}
	leftSubSum := int(math.MinInt16)
	curSubSum := 0
	for i := p; i > left-1; i-- {
		curSubSum += nums[i]
		leftSubSum = int(math.Max(float64(leftSubSum), float64(curSubSum)))
	}
	rightSubSum := int(math.MinInt16)
	curSubSum = 0
	for i := p + 1; i < right+1; i++ {
		curSubSum += nums[i]
		rightSubSum = int(math.Max(float64(rightSubSum), float64(curSubSum)))
	}

	return leftSubSum + rightSubSum
}

//MaxOf all numbers
func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
