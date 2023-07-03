package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

// X Loose
// Y Draw
// Z Win


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
        elk,roundResult := runeText[0], runeText[2]
        

        player := getPlayer(elk, roundResult)
        roundPoints := getPoints(roundResult) 

        currentPoints += (roundPoints + items[player])
    }

    fmt.Println(currentPoints)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// Not efficient
func getPlayer(elk rune, result rune) rune {
    switch {
        case result == 'Y': 
            return elk
        case result == 'Z':

            switch elk {
                case 'A':
                    return 'Y'
                case 'B':
                    return 'Z'
                default:
                    return 'X'
            }

        default: 
            switch elk {
                case 'A':
                    return 'Z'
                case 'B':
                    return 'X'
                default:
                    return 'Y'
            }
    }
}


func getPoints(result rune) int {
    m := map[rune]int{
        'Y': 3, // Draw
        'Z': 6, // Win
        'X': 0, // Loose
    }
    return m[result]
}
