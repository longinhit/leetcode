/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (34.74%)
 * Total Accepted:    185.5K
 * Total Submissions: 533.5K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 *
 *
 *
 *
 *
 */

// package main

// import "fmt"

// func main() {

// 	fmt.Printf("%v\n", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
// }

func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 1; j <= k; j++ {
			if i+j < length && nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}
