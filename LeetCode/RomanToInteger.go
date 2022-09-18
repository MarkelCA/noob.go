package main

import (
    "fmt"
)

type roman string

func main() {

    // Prefixing I example
    fmt.Println( romanToInt("LXIXV") )
    // Prefixing X example
    fmt.Println( romanToInt("CXLV") )
    // Prefixing C example
    fmt.Println( romanToInt("MCDI") )
}

var romanSymbols = map[int32]int {
    'I' : 1,
    'V' : 5,
    'X' : 10,
    'L' : 50,
    'C' : 100,
    'D' : 500,
    'M' : 1000,
}
func romanToInt(s string) (result int) {

    for i := 0 ; i < len(s) ; i++ {
        symbol := int32(s[i])

        r := roman(s)
        if r.isNotLastLetter(i) && r.isPrefixRune(symbol, i) {
            minus := r.getMinus(symbol, i)
            result +=  minus
            i+=1 // We skip the next rune because we already sum it
       } else {
            result += romanSymbols[ symbol ]
        }
    }


    return result
}

func (s roman) isNotLastLetter(i int) bool {
    return i != len(s)-1
}

func (s roman) isPrefixRune(char int32, index int) bool {
    possiblePrefixValues := [3]int32{'I', 'X', 'C'} 

    if !sliceContains(possiblePrefixValues[:], char) {
        return false
    }

    nextRune := int32(s[index+1])

    return isPrefixI(char, nextRune) || isPrefixX(char, nextRune) || isPrefixC(char, nextRune)
}

func sliceContains(slice []int32, char int32) bool {
    for _, v := range slice {
        if v == char {
            return true
        }
    }
    return false
}

func (s roman) getMinus(symbol int32, index int) int {
    nextRune := int32(s[index+1])
    return romanSymbols[nextRune] - romanSymbols[symbol]
}

func isPrefixI(char int32, nextRune int32) bool {
    prefixedPossibleValues := []int32{'V', 'X'}
    return isPrefix(char, 'I', nextRune,prefixedPossibleValues)
}

func isPrefixX(char int32, nextRune int32) bool {
    prefixedPossibleValues := []int32{'L', 'C'}
    return isPrefix(char, 'X', nextRune,prefixedPossibleValues)
}

func isPrefixC(char int32, nextRune int32) bool {
    prefixedPossibleValues := []int32{'D', 'M'}
    return isPrefix(char, 'C', nextRune,prefixedPossibleValues)
}

func isPrefix(char int32, expectedChar int32, nextRune int32, prefixedPossibleValues []int32) bool {
    if char != expectedChar {
        return false
    }
    return sliceContains(prefixedPossibleValues, nextRune)
}

