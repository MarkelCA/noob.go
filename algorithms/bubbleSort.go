package main

import("fmt")

func bubbleSort(ar []int) []int {
    for i,_ := range ar {
        for j:=0 ; j < len(ar) - i - 1 ; j++ {
            if ar[j] > ar[j+1] {
                ar[j], ar[j+1] = ar[j+1], ar[j]
            }
        }
    }

    return ar
}

func main() {
    ar := []int{23,41,25,54,18,14}
    bubbleSort(ar)
    fmt.Println(ar)
}
