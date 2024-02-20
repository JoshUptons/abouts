a = "11"
b = "1"


def addBinary(a: str, b: str) -> str:
    return f"{int(a, 2) + int(b, 2):b}"


print(f"{a} + {b} = {addBinary(a, b)}")
a = "10000"
b = "111"
print(f"{a} + {b} = {addBinary(a, b)}")
