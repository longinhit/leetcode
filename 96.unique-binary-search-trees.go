/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (45.76%)
 * Total Accepted:    193.5K
 * Total Submissions: 422.9K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */

/*
 * F(i,n) 表示n个节点，以i为根节点的搜索二叉树的种类数量。
 * G(n) 表示n个节点的二叉搜索树的种类数量
 * F(i,n) = G(i-1)*G(n-i)
 * G(n) = F(1,n) + F(2,n) + F(3,n) + ... + F(n-i,n) + ... + F(n-1,n)
 * G(n) = G(0)*G(n-1) + G(1)*G(n-2) + G(2)*G(n-3) + ... + G(i-1)*G(n-i) + ... + G(n-1)*G(0)
 * G（0) = G(1) = 1
 */
func numTrees(n int) int {
	g := []int{1, 1}
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += (g[j] * g[i-j-1])
		}
		g = append(g, sum)
	}
	return g[len(g)-1]
}

