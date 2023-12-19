
/* Reverse a string*/

// with array reverse
function reverseString(str) {
    return [...str].reverse().join("")
}

// with recursion
function recursiveReverseString(str) {
    if (str.length === 0) return ""

    return str[str.length-1] + 
        recursiveReverseString(str.slice(0, str.length-1))
}

console.log(reverseString("abcd"))
console.log(reverseString("1234"))

console.log(recursiveReverseString("abcd"))
