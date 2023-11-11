const mergesort = (list) => {
    if (list.length <= 1) return list

    let left = [], right = []
    for (let i = 0; i < list.length; i++) {
        (i < list.length / 2) ? left.push(list[i]) : right.push(list[i])
    }
    left = mergesort(left)
    right = mergesort(right)
    return merge(left, right)
}

const merge = (left, right) => {
    let result = []
    while (left.length > 0 && right.length > 0) {
        let a = left[0]
        let b = right[0]
        if (a <= b) {
            result.push(a)
            left.shift()
        } else {
            result.push(b)
            right.shift()
        }
    }
    while (left.length > 0) {
        result.push(left.shift())
    }
    while (right.length > 0) {
        result.push(right.shift())
    }
    return result
}

let arr = mergesort([1, 9, 53, 3, 8, 4, 8, 2, 5, 2, 23, 12, 1, 9, 2, 32, 3, 12, 7])
console.log(arr)
