def strStr(needle: str, haystack: str):
    l = 0
    r = len(needle)
    if haystack[l:r] == needle:
        return l

    while r < len(haystack):
        l += 1
        r += 1
        cur = haystack[l:r]
        if cur == needle:
            return l

    return -1


print(
    f"haystack='this is a test', needle='test', result={strStr('test', 'this is a test')}"
)

print(
    f"haystack='leetcode', needle='leeto', result={strStr('leeto', 'leetcode')}"
)
