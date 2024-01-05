package main

import "fmt"

func main() {
	strs := []string{
		"bccd",
		"a",
		"",
		"abradkadabra",
		"acad",
	}

	for _, s := range strs {
		fmt.Println(s, longestPalindrome(s))
	}
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}
	index := 0
	resLength := 0
	resIndices := make([]int, 2)

	// odd number palindromes
	for index <= len(s)-1 {
		p1 := index - 1
		p2 := index + 1
		for p1 >= 0 && p2 <= len(s)-1 && s[p1] == s[p2] {
			length := p2 - p1 + 1
			if length > resLength {
				resLength = length
				resIndices[0] = p1
				resIndices[1] = p2
			}
			p1--
			p2++
		}
		index++
	}

	// even number palindromes
	index = 0
	for index <= len(s)-1 {
		p1 := index
		p2 := index + 1
		for p1 >= 0 && p2 <= len(s)-1 && s[p1] == s[p2] {
			length := p2 - p1 + 1
			if length > resLength {
				resLength = length
				resIndices[0] = p1
				resIndices[1] = p2
			}
			p2++
			p1--
		}
		index++
	}

	return s[resIndices[0] : resIndices[1]+1]
}
