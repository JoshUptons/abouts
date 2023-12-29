package main

import "fmt"

// solution for leetcode #3 medium problem

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abbcaa"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	p1 := 0
	p2 := 1
	max := 1
	cur := 1
	seen := map[byte]bool{}
	seen[s[0]] = true
	for p2 < len(s) {
		if seen[s[p2]] {
			if p2 == len(s)-1 {
				return max
			}
			for s[p1] != s[p2] {
				seen[s[p1]] = false
				p1++
				cur--
			}
			p1++
		} else {
			seen[s[p2]] = true
			cur++
		}
		if cur > max {
			max = cur
		}
		p2++
	}
	return max
}
