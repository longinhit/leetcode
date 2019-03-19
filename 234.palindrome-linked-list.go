/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (35.38%)
 * Total Accepted:    237.1K
 * Total Submissions: 669K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */

// package main

// import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {
// 	head := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 1,
// 			Next: &ListNode{
// 				Val: 2,
// 			},
// 		},
// 	}
// 	fmt.Printf("%v\n", isPalindrome(head))
// }

func isPalindrome(head *ListNode) bool {
	length := 0
	node := head
	stack := []int{}
	for node != nil {
		length++
		node = node.Next
	}
	if length == 0 || length == 1 {
		return true
	}
	node = head
	count := 0
	if length%2 == 0 {
		for node != nil {
			count++
			if count <= length/2 {
				stack = append(stack, node.Val)
			} else {
				if stack[len(stack)-1] != node.Val {
					return false
				}
				stack = stack[0 : len(stack)-1]
			}
			node = node.Next
		}
	} else {
		for node != nil {
			count++
			if count <= length/2 {
				stack = append(stack, node.Val)
			} else if count > length/2+1 {
				if stack[len(stack)-1] != node.Val {
					return false
				}
				stack = stack[0 : len(stack)-1]
			}
			node = node.Next
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
