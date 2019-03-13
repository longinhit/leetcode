/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (35.31%)
 * Total Accepted:    209.1K
 * Total Submissions: 591.9K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
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
// 				Val: 6,
// 				Next: &ListNode{
// 					Val: 3,
// 					Next: &ListNode{
// 						Val: 4,
// 						Next: &ListNode{
// 							Val: 5,
// 							Next: &ListNode{
// 								Val: 6,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	fmt.Printf("dddddd\n")
// 	newHead := removeElements(head, 6)
// 	// newHead := head
// 	printLinks(newHead)
// }

// func printLinks(head *ListNode) *ListNode {
// 	node := head
// 	for {
// 		if node != nil {
// 			fmt.Printf("%d----->", node.Val)
// 			node = node.Next
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Printf("\n")
// 	return head
// }

func removeElements(head *ListNode, val int) *ListNode {
	node := head
	var preNode *ListNode
	for {
		if node != nil {
			if node.Val == val {
				if node.Next == nil {
					if preNode == nil {
						return nil
					}
					preNode.Next = nil
					node = nil
				} else {
					node.Val = node.Next.Val
					node.Next = node.Next.Next
				}
			} else {
				preNode = node
				node = node.Next
			}
		} else {
			break
		}
	}
	return head
}
