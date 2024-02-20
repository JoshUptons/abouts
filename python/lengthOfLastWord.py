"""
return the length of the last word that does not consist of empty spaces
so "hello world " would return 5, because you exclude the right empty space
"""


def lengthOfLastWord(s: str) -> int:
    length = 0

    if len(s) == 0:
        return 0

    while len(s) > 0 and s[-1] == ' ':
        s = s[:-1]

    if len(s) == 0:
        return 0

    while len(s) > 0 and s[-1] != ' ':
        length += 1
        s = s[:-1]

    return length


base = "hello world   "
base2 = "hello world"
base3 = "t s"
print(lengthOfLastWord(base))
print(lengthOfLastWord(base2))
print(lengthOfLastWord(base3))
