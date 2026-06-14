package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func longestConsecutive(nums []int) (res int) {
	cmp := map[int]int{}
	for i := range nums {
		cmp[nums[i]]++
	}
	for i := range cmp {
		k := i
		if cmp[k-1] > 0 {
			continue
		}
		for cmp[k] > 0 {
			k++
		}
		res = max(res, k-i)
	}
	return
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 || i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

func lengthOfLongestSubstring(s string) (ans int) {
	cmp := map[byte]int{}
	var left int
	for i := range s {
		cmp[s[i]]++
		for cmp[s[i]] > 1 {
			cmp[s[left]]--
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	if n <= 1 {
		return f[0]
	}
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func trap(height []int) (ans int) {
	n := len(height)
	lmx, rmx := make([]int, n), make([]int, n)
	l, r := height[0], height[n-1]
	for i := range height {
		l = max(height[i], l)
		r = max(height[n-1-i], r)
		lmx[i], rmx[n-1-i] = l, r
	}
	for i := range height {
		ans += min(lmx[i], rmx[i]) - height[i]
	}
	return
}

func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{}
	sum := 0
	mp[0] = 1
	for i := range nums {
		sum += nums[i]
		ans += mp[sum-k]
		mp[sum]++
	}

	return
}

func lengthOfLIS(nums []int) int {
	f := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		var temp int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp = max(f[j+1], temp)
			}
		}
		f[i+1] = temp + 1
	}
	return slices.Max(f[1:])
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt / 2
	}
	for i := 1; i <= amount; i++ {
		for j := range coins {
			if coins[j] > i {
				continue
			}
			dp[i] = min(dp[i-coins[j]]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt/2 {
		return -1
	}
	return dp[amount]
}

func canPartition(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	for i := range nums {
		for j := target; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
		}
	}
	return f[target]
}

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}
	used := make([]bool, n)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(idx + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}

func subsets(nums []int) (ans [][]int) {
	subset := []int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, subset...))
			return
		}
		dfs(idx + 1)
		subset = append(subset, nums[idx])
		dfs(idx + 1)
		subset = subset[:len(subset)-1]
	}
	dfs(0)
	return
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	list := []int{}
	var findTarget func(idx int, num int)
	findTarget = func(idx int, num int) {
		if num == 0 {
			ans = append(ans, append([]int{}, list...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] > num {
				break
			}
			list = append(list, candidates[i])
			findTarget(i, num-candidates[i])
			list = list[:len(list)-1]
		}
	}
	findTarget(0, target)
	return
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[n1][n2]
}

func numberOfArithmeticSlices(nums []int) (ans int) {
	if len(nums) <= 2 {
		return
	}
	d, t := nums[0]-nums[1], 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return
}

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var bfs func(i, j int)
	bfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		bfs(i+1, j)
		bfs(i, j+1)
		bfs(i-1, j)
		bfs(i, j-1)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				bfs(i, j)
			}
		}
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	queue := []int{}
	for i := range nums {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i-queue[0] >= k {
			queue = queue[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return
}

func main() {
	res := longestConsecutive([]int{1, 0, 1, 2})
	fmt.Println(res)
	fmt.Println("woojy")
	fmt.Println("woojy2")
}
