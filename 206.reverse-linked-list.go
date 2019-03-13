/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (52.92%)
 * Total Accepted:    526.4K
 * Total Submissions: 993.4K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
 *
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val: 5,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	newHead := reverseList(head)
// 	printList(newHead)
// }

// func printList(head *ListNode) {
// 	node := head
// 	for {
// 		if node == nil {
// 			break
// 		}
// 		fmt.Printf("%d--->", node.Val)
// 		node = node.Next
// 	}
// 	fmt.Printf("\n")
// 	return
// }

func reverseList(head *ListNode) *ListNode {
	node := head
	var newHead *ListNode
	for {
		if node == nil {
			break
		}
		if newHead == nil {
			newHead = node
			node = node.Next
			newHead.Next = nil
		} else {
			temp := node
			node = node.Next
			temp.Next = newHead
			newHead = temp
		}
	}
	return newHead
}

