package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("122", reverse(122))
	fmt.Println("4331", reverse(4331))
	fmt.Println("-211", reverse(-211))
	fmt.Println("0", reverse(0))
	fmt.Println("120", reverse(120))
	fmt.Println("12345678912", reverse(12345678912))
}

func reverse(x int) int {
	negative := x < 0
	if negative {
		x *= -1
	}
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + (x % 10)
		x = int(math.Floor(float64(x / 10)))
	}
	if reversed > 1<<31 {
		return 0
	}
	if negative {
		reversed *= -1
	}
	return reversed
}
