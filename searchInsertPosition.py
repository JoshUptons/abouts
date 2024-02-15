from typing import List


def bs(list: List[int], high: int, low: int, target: int) -> int:
    m = (high + low) // 2
    if m == 0 and list[0] != target:
        return 0
    if m == len(list) - 1 and list[len(list) - 1] != target:
        return len(list)
    if target == list[m]:
        return m
    elif target < list[m]:
        if target > list[m - 1]:
            return m
        elif target == list[m - 1]:
            return m - 1
        else:
            return bs(list, m, 0, target)
    else:
        if target < list[m + 1] or target == list[m + 1]:
            return m + 1
        else:
            return bs(list, len(list), m, target)


def searchInsert(nums: List[int], target: int) -> int:
    return bs(nums, len(nums), 0, target)


l = [1, 2, 4, 9, 12, 24]
result = searchInsert(l, 0)
print(result)