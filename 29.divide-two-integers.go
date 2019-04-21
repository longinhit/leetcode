/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.15%)
 * Total Accepted:    189.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 *
 * Return the quotient after dividing dividend by divisor.
 *
 * The integer division should truncate toward zero.
 *
 * Example 1:
 *
 *
 * Input: dividend = 10, divisor = 3
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: dividend = 7, divisor = -3
 * Output: -2
 *
 * Note:
 *
 *
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 *
 *
 */

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%v\n", divide(6, 1))
// }
func divide(dividend int, divisor int) int {
	result := 0
	flag := 0
	if divisor < 0 {
		divisor = 0 - divisor
		flag = 1
	}
	if dividend < 0 {
		flag ^= 1
		dividend = 0 - dividend
	}

	for dividend >= divisor {
		quotientTmp := 1
		divisorTmp := divisor
		for dividend > divisorTmp<<1 {
			quotientTmp = quotientTmp << 1
			divisorTmp = divisorTmp << 1
		}
		dividend -= divisorTmp
		result += quotientTmp
	}
	if flag == 1 {
		return 0 - result
	}
	if result > (1<<31)-1 {
		return (1 << 31) - 1
	}
	return result
}
