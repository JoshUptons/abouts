package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(betterZigZag("PAYPALISHIRING", 3))
	fmt.Println(betterZigZag("PAYPALISHIRING", 4))
	fmt.Println(betterZigZag("PAYPALISHIRING", 2))
}

func zigzag(str string, rows int) string {
	grid := make([][]string, rows)

	row := 0
	dir := 1

	for i, s := range str {

		grid[row] = append(grid[row], string(s))

		row += dir
		// if the row is at the upper bound
		if row == rows-1 {
			dir = -dir
		}
		// if the row is at the lower bound and not the first character
		if row == 0 && i > 0 {
			dir = -dir
		}
	}

	fmt.Println(grid)

	var output string
	for i := 0; i < len(grid); i++ {
		output += strings.Join(grid[i], "")
	}

	return output
}

// after looking through some examples, I ran across the strings.Builder, which
// I have never used before, this would be a lower memory usage solution as you
// writing to a buffer rather than storing the grid on the heap.

func betterZigZag(str string, rows int) string {
	if rows <= 1 {
		return str
	}
	grid, dir := make([]strings.Builder, rows), 1

	for i, j := 0, 0; i < len(str); i++ {
		grid[j].WriteByte(str[i])
		j += dir
		if j == rows-1 || j == 0 {
			dir *= -1
		}
	}

	var result strings.Builder
	for i := 0; i < len(grid); i++ {
		result.WriteString(grid[i].String())
	}
	return result.String()
}
