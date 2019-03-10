/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (42.79%)
 * Total Accepted:    466.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */
// package main

// import "fmt"

// func main() {
// 	fmt.Printf("\n%d\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
// }

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sumTemp := 0
	max := int(-1) << 63
	for _, num := range nums {
		if sumTemp+num > num {
			sumTemp += num
		} else {
			sumTemp = num
		}
		if sumTemp > max {
			max = sumTemp
		}
	}
	return max
}
