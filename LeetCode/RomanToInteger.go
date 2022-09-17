package main

import (
    "fmt"
    "golang.org/x/exp/slices"
)

type roman string

func main() {
    //fmt.Println( romanToInt("MCMXCIV") )
    //fmt.Println( romanToInt("XIVI") )
    fmt.Println( romanToInt("LXIXV") )
    //fmt.Println( romanToInt("MCMIXCV") )
    //fmt.Println( roman2("MCMXCIV") )
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


func (s roman) isNotLastLetter(i int) bool {
    return i != len(s)-1
}

func isMinusRune(char int32) bool {
    possibleValues := [3]int32{'I', 'X', 'C'} 
    return slices.Contains(possibleValues[:], char)
}

func (s roman) getMinus(symbol int32, index int) int {
    nextRune := int32(s[index+1])
    return romanSymbols[nextRune] - romanSymbols[symbol]
}

func romanToInt(s roman) (result int) {

    for i := 0 ; i < len(s) ; i++ {
        symbol := int32(s[i])

            if isMinusRune(symbol) && s.isNotLastLetter(i) {
            nextRune := int32(s[i+1])


            if symbol == 'I' && ( nextRune == 'V' || nextRune == 'X') {
                result += romanSymbols[nextRune] - 1 // We sum the next rune minus one
                i+=1 // We skip the next rune
            } else {
                result += romanSymbols[ symbol ]
            }
            //if isMinusRune(symbol) {
                //minus := s.getMinus(symbol, i)
                //fmt.Println(minus)
                //result +=  minus
            //}
       } else {
            result += romanSymbols[ symbol ]
        }
    }



    return result
}
