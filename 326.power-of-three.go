/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 *
 * https://leetcode.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (41.49%)
 * Total Accepted:    175.9K
 * Total Submissions: 423.8K
 * Testcase Example:  '27'
 *
 * Given an integer, write a function to determine if it is a power of three.
 *
 * Example 1:
 *
 *
 * Input: 27
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: 0
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: 9
 * Output: true
 *
 * Example 4:
 *
 *
 * Input: 45
 * Output: false
 *
 * Follow up:
 * Could you do it without using any loop / recursion?
 */
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	} else if n == 0 || n%3 != 0 {
		return false
	} else {
		return isPowerOfThree(n / 3)
	}
}

