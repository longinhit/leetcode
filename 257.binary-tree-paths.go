/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (45.09%)
 * Total Accepted:    213.8K
 * Total Submissions: 472.9K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 *
 */

// package main

import "fmt"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func main() {
// 	fmt.Printf("%d", 45)
// }

var allLeafPaths []string

func printSilce(pathSlice []int) string {
	pathStr := ""
	for _, path := range pathSlice {
		if len(pathStr) == 0 {
			pathStr = fmt.Sprintf("%d", path)
		} else {
			pathStr = fmt.Sprintf("%s->%d", pathStr, path)
		}
	}
	return pathStr
}

func printTreePath(root *TreeNode, pathSlice []int) {
	if root == nil {
		return
	}
	pathSlice = append(pathSlice, root.Val)
	if root.Left == nil && root.Right == nil {
		if len(pathSlice) > 0 {
			allLeafPaths = append(allLeafPaths, printSilce(pathSlice))
		}
	}
	if root.Left != nil {
		printTreePath(root.Left, pathSlice)
	}
	if root.Right != nil {
		printTreePath(root.Right, pathSlice)
	}
	pathSlice = pathSlice[0 : len(pathSlice)-1]
	return
}

func binaryTreePaths(root *TreeNode) []string {
	allLeafPaths = []string{}
	if root == nil {
		return allLeafPaths
	}
	printTreePath(root, []int{})
	return allLeafPaths
}
