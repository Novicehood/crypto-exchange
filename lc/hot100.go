package main

import (
	"fmt"
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

func main() {
	res := longestConsecutive([]int{1, 0, 1, 2})
	fmt.Println(res)
}
