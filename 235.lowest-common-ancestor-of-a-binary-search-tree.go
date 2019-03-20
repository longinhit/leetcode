/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (43.55%)
 * Total Accepted:    262.4K
 * Total Submissions: 600.8K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * Given a binary search tree (BST), find the lowest common ancestor (LCA) of
 * two given nodes in the BST.
 *
 * According to the definition of LCA on Wikipedia: “The lowest common ancestor
 * is defined between two nodes p and q as the lowest node in T that has both p
 * and q as descendants (where we allow a node to be a descendant of itself).”
 *
 * Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * Output: 6
 * Explanation: The LCA of nodes 2 and 8 is 6.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * Output: 2
 * Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant
 * of itself according to the LCA definition.
 *
 *
 *
 *
 * Note:
 *
 *
 * All of the nodes' values will be unique.
 * p and q are different and both values will exist in the BST.
 *
 *
 */

// package main

// import "fmt"

// func main() {

// 	root := &TreeNode{
// 		Val: 2,
// 		Left: &TreeNode{
// 			Val: 1,
// 			// Left: &TreeNode{
// 			// 	Val: 0,
// 			// },
// 			// Right: &TreeNode{
// 			// 	Val: 4,
// 			// 	Left: &TreeNode{
// 			// 		Val: 3,
// 			// 	},
// 			// 	Right: &TreeNode{
// 			// 		Val: 5,
// 			// 	},
// 			// },
// 		},
// 		// Right: &TreeNode{
// 		// 	Val: 8,
// 		// 	Left: &TreeNode{
// 		// 		Val: 7,
// 		// 	},
// 		// 	Right: &TreeNode{
// 		// 		Val: 9,
// 		// 	},
// 		// },
// 	}

// 	fmt.Printf("%+v\n", lowestCommonAncestor(root, &TreeNode{
// 		Val: 1,
// 	}, &TreeNode{
// 		Val: 2,
// 	}))
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func printPath(path []*TreeNode) {
// 	for _, node := range path {
// 		fmt.Printf("%d-->", node.Val)
// 	}
// 	fmt.Printf("\n")
// }

func preOrder(root *TreeNode, target *TreeNode, path []*TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target.Val {
		path = append(path, root)
		return path
	}
	if root.Left != nil {
		path = append(path, root)
		newPath := preOrder(root.Left, target, path)
		if newPath != nil {
			return newPath
		}
		path = path[0 : len(path)-1]
	}
	if root.Right != nil {
		path = append(path, root)
		newPath := preOrder(root.Right, target, path)
		if newPath != nil {
			return newPath
		}
		path = path[0 : len(path)-1]
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := preOrder(root, p, []*TreeNode{})
	qPath := preOrder(root, q, []*TreeNode{})
	// printPath(pPath)
	// printPath(qPath)
	if len(pPath) > len(qPath) {
		for index, _ := range pPath {
			if index < len(qPath) {
				if pPath[index] != qPath[index] {
					return pPath[index-1]
				}
			} else if index == len(qPath) {
				return pPath[index-1]
			} else {
				break
			}
		}
	} else {
		for index, _ := range qPath {
			if index < len(pPath) {
				if pPath[index] != qPath[index] {
					return qPath[index-1]
				}
			} else if index == len(pPath) {
				return pPath[index-1]
			} else {
				break
			}
		}
	}
	return nil
}
