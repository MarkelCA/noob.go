package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

var items map[rune]int = map[rune]int{
    'A': 1, // Rock
    'B': 2, // Paper
    'C': 3, // Sciccors

    'X': 1, // Rock
    'Y': 2, // Paper
    'Z': 3, // Sciccors
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

    currentPoints := 0

    scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

        runeText := []rune(text)
        elk,player := runeText[0], runeText[2]

        roundPoints := move(elk,player)
        currentPoints += (roundPoints + items[player])
    }

    fmt.Println(currentPoints)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func move(elk rune, player rune) int {
    switch {
        case items[elk] == items[player]: // Draw
            return 3
        case  // Win
            items[player] == 1 && items[elk] == 3,
            items[player] == 2 && items[elk] == 1,
            items[player] == 3 && items[elk] == 2:
            return 6
        default: // Lose
            return 0
    }
}
