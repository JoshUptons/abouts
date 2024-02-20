"""
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

    Input: n = 2
    Output: 2
    Explanation: There are two ways to climb to the top.
    1. 1 step + 1 step
    2. 2 steps

Example 2:

    Input: n = 3
    Output: 3

    Explanation: There are three ways to climb to the top.
    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step
 

Constraints:

    1 <= n <= 45


we can start at the constraint the one of the ways will always be 1 * n steps
then we need to work backwards from there, determining each combo of steps including
reordering of the steps taken.

if there are 5 steps
the steps can be an array like [5,3,2,1,1]
where there are 5 ways to get to the 5th step
"""


def climbStairs(n: int) -> int:
    one, two = 1, 1
    for i in range(n - 1):
        temp = one
        one = one + two
        two = temp

    return one


print(climbStairs(5))
print(climbStairs(0))
print(climbStairs(2))
print(climbStairs(33))
print(climbStairs(40))
