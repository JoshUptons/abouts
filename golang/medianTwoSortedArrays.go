package main

import "fmt"

// solution for the leetcode #4 median of two sorted arrays, hard problem.

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	med := findMedianSortedArrays(nums1, nums2)
	fmt.Println("median: ", med) // expect 2
}

/*
	start a divider in the center of the smaller array
	and the other divider at the length of the first array
*/

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var n, m int
	n = len(nums1)
	m = len(nums2)

	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	s := 0
	e := n

	for s <= e {

		var d1, d2 int
		d1 = s + (e-s)/2
		d2 = (n+m+1)/2 - d1

		var l1, l2, r1, r2 int
		if d1-1 >= 0 {
			l1 = nums1[d1-1]
		} else {
			l1 = -1 << 31
		}

		if d1 < n {
			r1 = nums1[d1]
		} else {
			r1 = 1 << 31
		}

		if d2-1 >= 0 {
			l2 = nums2[d2-1]
		} else {
			l2 = -1 << 31
		}

		if d2 < m {
			r2 = nums2[d2]
		} else {
			r2 = 1 << 31
		}

		if l1 <= r2 && l2 <= r1 {
			var median float64

			if (n+m)%2 == 1 {
				// is odd
				if l1 > l2 {
					median = float64(l1)
				} else {
					median = float64(l2)
				}
			} else {
				// is even
				var maxLeft, minRight float64
				if l1 > l2 {
					maxLeft = float64(l1)
				} else {
					maxLeft = float64(l2)
				}

				if r1 < r2 {
					minRight = float64(r1)
				} else {
					minRight = float64(r2)
				}

				median = (maxLeft + minRight) / 2.0
			}

			return median
		}

		if l1 > r2 {
			// shift divider left
			e = d1 - 1
		}

		if l2 > r1 {
			// shift diveider right
			s = d1 + 1
		}

	}

	return 0.0
}
