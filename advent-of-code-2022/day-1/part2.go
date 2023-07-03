package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	currentCals := 0
	currentElk := 1
	topThree := [3]int{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {

			if currentCals > topThree[2] {
				if currentCals > topThree[1] {
					if currentCals > topThree[0] {
						swap(topThree[0:], currentCals)
						topThree[0] = currentCals
					} else {
						swap(topThree[1:], currentCals)
						topThree[1] = currentCals
					}
				} else {
					topThree[2] = currentCals
				}
			}
			currentElk += 1
			currentCals = 0

		} else {
			i, _ := strconv.Atoi(text)
			currentCals += i
		}

	}

	fmt.Println(topThree)

	sum := 0
	for _, x := range topThree {
		sum += x
	}

	fmt.Printf("Top three sum: %v\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Sets the passed value to slice[0] position
// and pushes the rest of the values one position
// to the right.
func swap(slice []int, val int) {
	aux := 0
	for i, x := range slice {
		aux = x
		slice[i] = val
		val = aux
	}
}
