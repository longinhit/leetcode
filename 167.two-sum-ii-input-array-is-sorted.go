/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (49.31%)
 * Total Accepted:    213.1K
 * Total Submissions: 432.2K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers that is already sorted in ascending order, find
 * two numbers such that they add up to a specific target number.
 *
 * The function twoSum should return indices of the two numbers such that they
 * add up to the target, where index1 must be less than index2.
 *
 * Note:
 *
 *
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 *
 *
 * Example:
 *
 *
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 *
 */
// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))

// }

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return []int{}
	}
	start := 0
	end := len(numbers) - 1
	for start <= end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			break
		} else if sum > target {
			end -= 1
		} else {
			start += 1
		}
	}
	return []int{start + 1, end + 1}
}
