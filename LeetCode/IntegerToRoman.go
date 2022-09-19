package main

import ("fmt" ; "strings")

type roman struct {
    Letter   string
    Value int
}
var runes = []roman {
    roman{"M" ,  1000 },
    roman{"CM",  900 },
    roman{"D" ,  500 },
    roman{"CD",  400 },
    roman{"C" ,  100 },
    roman{"XC",  90 },
    roman{"L" ,  50 },
    roman{"XL",  40 },
    roman{"X" ,  10 },
    roman{"IX",  9  },
    roman{"V" ,  5  },
    roman{"IV",  4  },
    roman{"I" ,  1  },
}



func main() {
    n := 20
    fmt.Println( intToRoman(n) )
} 

func intToRoman(num int) (result string) {
    for num > 0 {
        for _,r := range runes {
            if num >= r.Value {
                times := num / r.Value
                val   := times * r.Value

                result += strings.Repeat(r.Letter, times)
                num    -= val
            }
        }
    }

    return result
}
