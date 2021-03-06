/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (41.66%)
 * Total Accepted:    215.5K
 * Total Submissions: 517.1K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: true
 * Explanation: 2^0 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 *
 * Example 3:
 *
 *
 * Input: 218
 * Output: false
 *
 */

// package main

// import "fmt"

// func main() {
// 	totalTables := 20
// 	fmt.Printf(fmt.Sprintf("%s_%01d", "tablePrefix", int(19)%totalTables))
// }

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	tmp := n

	for {
		if tmp == 1 {
			return true
		}
		if tmp%2 != 0 {
			return false
		}
		tmp /= 2
	}
}
