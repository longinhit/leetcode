/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 *
 * https://leetcode.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (34.17%)
 * Total Accepted:    140.7K
 * Total Submissions: 409.4K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, return
 * the ordering of courses you should take to finish all courses.
 *
 * There may be multiple correct orders, you just need to return one of them.
 * If it is impossible to finish all courses, return an empty array.
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished
 * course 0. So the correct course order is [0,1] .
 *
 * Example 2:
 *
 *
 * Input: 4, [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,1,2,3] or [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both
 * ⁠            courses 1 and 2. Both courses 1 and 2 should be taken after you
 * finished course 0.
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3] .
 *
 * Note:
 *
 *
 * The input prerequisites is a graph represented by a list of edges, not
 * adjacency matrices. Read more about how a graph is represented.
 * You may assume that there are no duplicate edges in the input
 * prerequisites.
 *
 *
 */

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("%v", findOrder(2, [][]int{{0, 1}}))
// }
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := [][]int{}
	vistor := map[int]int{}
	for i := 0; i < numCourses; i++ {
		temp := []int{}
		for j := 0; j < numCourses; j++ {
			temp = append(temp, 0)
		}
		graph = append(graph, temp)
	}
	result := []int{}
	for _, edge := range prerequisites {
		graph[edge[0]][edge[1]] = 1
	}
	for len(result) < numCourses {
		hasPrerequisit := 1
		for i := 0; i < numCourses; i++ {
			if _, ok := vistor[i]; ok {
				continue
			}
			indegree := 0
			for j := 0; j < numCourses; j++ {
				if graph[i][j] == 1 {
					indegree++
				}
			}
			if indegree == 0 {
				vistor[i] = i
				result = append(result, i)
				for j := 0; j < numCourses; j++ {
					graph[j][i] = 0
				}
				hasPrerequisit = 0
			}
		}
		if hasPrerequisit == 1 {
			return []int{}
		}
	}
	return result
}
