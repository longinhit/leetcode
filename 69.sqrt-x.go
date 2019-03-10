/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (30.67%)
 * Total Accepted:    330.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("%d\n", mySqrt(4))
// }

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	start := 0
	end := x
	xTemp := x
	for start < end {
		mid := (start + end) / 2
		sqrtTemp := mid * mid
		if sqrtTemp > xTemp {
			end = mid
		} else if sqrtTemp < xTemp {
			start = mid + 1
		} else {
			return mid
		}
		// fmt.Printf("start:%d, end:%d, sqrt:%d,  sqrtTemp:%d,  xTemp:%d,  diff:%d\n", start, end, sqrt, sqrtTemp, xTemp, sqrtTemp-xTemp)
	}
	// fmt.Printf("start:%d, end:%d, sqrt:%d\n", start, end, sqrt)
	return end - 1
}
