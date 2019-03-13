/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (36.77%)
 * Total Accepted:    189.5K
 * Total Submissions: 515.1K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 */
// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%v", isIsomorphic("egg", "add"))
// }
func isIsomorphic(sssss string, ttttt string) bool {
	sMaping := map[byte]byte{}
	tMaping := map[byte]byte{}
	if len(sssss) != len(ttttt) {
		return false
	}
	for i := 0; i < len(sssss); i++ {
		t, tOk := sMaping[sssss[i]]
		s, sOk := tMaping[ttttt[i]]
		if !tOk && !sOk {
			sMaping[sssss[i]] = ttttt[i]
			tMaping[ttttt[i]] = sssss[i]
		} else {
			if t != ttttt[i] && sssss[i] != s {
				return false
			}
		}
	}
	return true
}
