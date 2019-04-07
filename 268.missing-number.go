/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 *
 * https://leetcode.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (47.88%)
 * Total Accepted:    259.4K
 * Total Submissions: 541.8K
 * Testcase Example:  '[3,0,1]'
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 *
 * Example 1:
 *
 *
 * Input: [3,0,1]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 *
 *
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant extra space complexity?
 */

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%+v\n", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
// }
func missingNumber(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	sum_1 := length * (length + 1) / 2
	sum_2 := 0
	for _, num := range nums {
		sum_2 += num
	}
	return sum_1 - sum_2
}
