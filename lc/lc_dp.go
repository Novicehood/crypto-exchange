package main

import (
	"math"
	"sort"
)

func findTargetSumWays(nums []int, target int) int {
	var S int
	for i := range nums {
		S += nums[i]
	}
	S -= target
	if S < 0 || S%2 == 1 {
		return 0
	}
	p, n := S/2, len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, p+1)
	}
	f[0][0] = 1
	for i := range nums {
		for j := 0; j <= p; j++ {
			if j < nums[i] {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = f[i][j] + f[i][j-nums[i]] // 不选 + 选
			}
		}
	}
	return f[n][p]
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	for j := 1; j <= target; j++ {
		f[0][j] = math.MinInt
	}
	for i := range nums {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j-nums[i]]+1, f[i][j])
			}
		}
	}
	if f[n][target] < 0 {
		return -1
	}
	return f[n][target]
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i := range intervals {
		if intervals[i][0] > end {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			end = max(end, intervals[i][1])
		}
	}
	res = append(res, []int{start, end})
	return res
}

func jump(nums []int) int {
	var count int
	max_bound, right_bound := 0, 0
	for i := range nums[:len(nums)-1] {
		max_bound = max(max_bound, i+nums[i])
		if i == right_bound {
			count++
			right_bound = max_bound
		}
	}
	return count
}

func main() {

}
