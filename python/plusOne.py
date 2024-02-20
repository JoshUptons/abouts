from typing import List
"""
given a list of numbers which when joined create a large integer,
find the new list of numbers, which result from incrementing the 
largest number by 1.

this is similar to the calculate a linked list of reverse numbers
but is a little simpler.
We are going to have a "remainder" check to see if the number needs
to roll over.
"""


def plusOne(digits: List[int]) -> List[int]:

    right = len(digits) - 1
    n = digits[right] + 1
    carry = n // 10
    digits[right] = n % 10
    while carry > 0 and right > 0:
        right -= 1
        n = digits[right] + 1
        carry = n // 10
        digits[right] = n % 10

    if carry > 0 and right == 0:
        digits.insert(0, carry)

    return digits


print(f'expect: [1, 0]: {plusOne([9])}: {[1, 0] == plusOne([9])}')
print(
    f'expect [1, 2, 4]: {plusOne([1, 2, 3])}: {[1, 2, 4] == plusOne([1, 2, 3])}'
)
print(
    f'expect [1, 0, 0, 0]: {plusOne([9, 9, 9])}: {[1, 0, 0, 0] == plusOne([9, 9, 9])}'
)
print(
    f'expect [2, 3, 4, 1]: {plusOne([2, 3, 4, 0])}: {[2, 3, 4, 1] == plusOne([2, 3, 4, 0])}'
)
