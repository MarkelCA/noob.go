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
	biggerCals := 0

	currentElk := 1
	richestElk := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			currentElk += 1

			if currentCals > biggerCals {
				richestElk = currentElk
				biggerCals = currentCals
			}

			currentCals = 0

		} else {
			i, _ := strconv.Atoi(text)
			currentCals += i
		}

	}

	fmt.Printf("Elk: %v, Calories: %v\n", richestElk, biggerCals)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
