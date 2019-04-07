/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (34.70%)
 * Total Accepted:    134.7K
 * Total Submissions: 388.3K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 *
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 *
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters that may be separated by a single space.
 *
 */
import "strings"

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Printf("%+v\n", wordPattern("abba", "dog dog dog dog"))
// }
func wordPattern(pattern string, str string) bool {
	charMapStr := map[rune]string{}
	strMapChar := map[string]rune{}
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	for index, ch1 := range pattern {
		subStr, ok := charMapStr[ch1]
		if ok {
			if subStr != strs[index] {
				return false
			}
		} else {
			if ch2, ok := strMapChar[strs[index]]; ok {
				if ch2 != ch1 {
					return false
				}
			}
			strMapChar[strs[index]] = ch1
			charMapStr[ch1] = strs[index]
		}
	}
	return true
}
