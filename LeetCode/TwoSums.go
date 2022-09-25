package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		} else {
			m[n] = i
		}
	}

	return nil
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
	//fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))
}
