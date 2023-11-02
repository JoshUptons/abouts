const bubblesort = (input) => {
    let p1 = 0
    let p2 = 1
    let allclear = true
    while (p1 <= input.length - 2) {
        if (input[p1] > input[p2]) {
            let swap = input[p1]
            input[p1] = input[p2]
            input[p2] = swap
            allclear = false
        }
        p1++
        p2++
    }
    return (allclear) ? input : bubblesort(input)
}

let arr = bubblesort([8, 3, 2, 6, 1, 5])
console.log(arr)
