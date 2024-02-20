"""
without using built in python functions, return the sqrt of an integer
rounded down to the nearest integer.

The square root is the smallest number that when squared, equals the given
number.
"""

from math import sqrt


def mySqrt(n: int) -> int:
    if n == 1:
        return 1
    if n == 0:
        return 0

    m = n // 2

    def bs(n: int, m: int) -> int:
        if m * m == n:
            return m
        elif m * m > n:
            return bs(n, m // 2)
        else:
            if (m + 1) * (m + 1) > n:
                return m
            return bs(n, m + 1)

    return bs(n, m)


print(mySqrt(12), sqrt(12))
print(mySqrt(4), sqrt(4))
print(mySqrt(124), sqrt(124))
print(mySqrt(1024), sqrt(1024))
print(mySqrt(0), sqrt(0))
print(mySqrt(1), sqrt(1))
print(mySqrt(2), sqrt(2))
print(mySqrt(390323), sqrt(390323))
