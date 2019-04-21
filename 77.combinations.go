/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (47.00%)
 * Total Accepted:    195.7K
 * Total Submissions: 416.3K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%+v\n", combine(4, 2))
// }

func getOneCombine(start, n, k int, oneCombine []int, result [][]int) [][]int {
	if k == 0 {
		temp := []int{}
		for _, item := range oneCombine {
			temp = append(temp, item)
		}
		result = append(result, temp)
	}
	for i := start; i <= n; i++ {
		oneCombine = append(oneCombine, i)
		result = getOneCombine(i+1, n, k-1, oneCombine, result)
		oneCombine = oneCombine[:len(oneCombine)-1]
	}
	return result
}

func combine(n int, k int) [][]int {
	return getOneCombine(1, n, k, []int{}, [][]int{})
}
