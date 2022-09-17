package main

import("fmt")

func Append(slice []int, nums ...int) []int {
    for _, num := range nums {
        slice = extend(slice, num)
    }
    return slice
}
func extend(slice []int, num int) []int {
    n := len(slice)
    if cap(slice) == n {
        newSlice := make([]int, n, 2*cap(slice))
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:n+1]
    slice[n] = num
    return slice
}

func main() {
    slice := make([]int, 3, 4)
    slice = extend(slice, 6)
    fmt.Println(slice)
    slice = extend(slice, 1)
    fmt.Println(slice)
    slice = extend(slice, 7)
    fmt.Println(slice)


    fmt.Println("...........")
    slice = Append(slice, 7,4,2,0)
    fmt.Println(slice)

    fmt.Println("...........")
    slice1 := []int{0, 1, 2, 3, 4}
    slice2 := []int{55, 66, 77}
    fmt.Println(slice1)
    slice1 = Append(slice1, slice2...) // The '...' is essential!
    fmt.Println(slice1)

}  
