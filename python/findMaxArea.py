from typing import List
"""
Problem, given an array of heights, length n
Find the largest area between heights that will hold water.
The area will be found as length x width which in this case
will be (x2 - x1) * (min(y1, y2))

4|     #
3| #   #           # 
2| #   #       #   #
1|_#___#___#___#___#____
0  0   1   2   3   4   5
    area will be between
"""


def findMaxArea(heights: List[int]) -> int:
    maxArea = 0
    p1 = 0
    p2 = len(heights) - 1
    while p1 < p2:
        area = (p2 - p1) * (min(heights[p1], heights[p2]))
        if area > maxArea:
            maxArea = area

        if heights[p1] < heights[p2]:
            p1 += 1
        else:
            p2 -= 1

    return maxArea


print(findMaxArea([2, 6, 3, 4, 1, 1, 5, 2]))
print(findMaxArea([3, 4, 5, 1, 2]))
