/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (51.04%)
 * Total Accepted:    310.5K
 * Total Submissions: 606.4K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 *
 * Example 1:
 *
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 *
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 *
 */
import (
	"reflect"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRune := []byte{}
	tRune := []byte{}
	for index := range s {
		sRune = append(sRune, s[index])
	}
	for index := range t {
		tRune = append(tRune, t[index])
	}
	sort.Slice(sRune, func(i, j int) bool {
		return sRune[i] < sRune[j]
	})
	sort.Slice(tRune, func(i, j int) bool {
		return tRune[i] < tRune[j]
	})
	return reflect.DeepEqual(sRune, tRune)
}
