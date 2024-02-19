from typing import List


def bs(nums: List[int], high: int, low: int, target: int) -> int:
    m = (high - low) // 2 + low
    if nums[m] == target:
        while nums[m - 1] == target:
            m -= 1
        return m
    if target < nums[m]:
        high = m
        if target > nums[m - 1] or m == 0:
            return m
    if target > nums[m]:
        low = m
        if (m < len(nums) - 1
                and target < nums[m + 1]) or (m == len(nums) - 1
                                              and target != nums[m]):
            return m + 1

    return bs(nums, high, low, target)


nums = [1, 3, 5, 7, 12, 15]
result = bs(nums, len(nums), 0, 16)

print(result)
