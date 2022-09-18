package main

import ("fmt" ; "sort")

var possibleRunes = map[string]int {
    "I"  : 1,
    "IV" : 4,
    "V"  : 5,
    "IX" : 9,
    "X"  : 10,
    "XL" : 40,
    "L"  : 50,
    "XC" : 90,
    "C"  : 100,
    "CD" : 400,
    "D"  : 500,
    "CM" : 900,
    "M"  : 1000,
}

type kv struct {
    Key   string
    Value int
}

func order(m map[string]int) (result []kv) {
    for k, v := range m {
        result = append(result, kv{k, v})
    }

    sort.Slice(result, func(i, j int) bool {
        return result[i].Value > result[j].Value
    })

    return result
}

func main() {
    n := 3
    fmt.Println( intToRoman(n) )
} 

func intToRoman(num int) (result string) {
    orderedRunes  := order(possibleRunes)

    for num > 0 {
        for _,o := range orderedRunes {
            if num >= o.Value {
                result += o.Key
                num -= o.Value
            }
        }
    }

    return result
    
}
