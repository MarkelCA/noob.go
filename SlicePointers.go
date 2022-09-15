// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"fmt"
)

type path []byte
type roman int

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		fmt.Printf("%p\n", (p))
		fmt.Printf("%p\n", *p)
		*p = (*p)[0:i]
		fmt.Printf("%p\n", *p)
	}
}

func (n roman) sum1() roman {
	return n + 1
}



func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
    fmt.Printf("%p\n", pathName)
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)

	var n roman = 1
	fmt.Println(n.sum1())

    pathName = path("/usr/bin/tso")
    fmt.Printf("%p\n", pathName)
    pathName.ToUpper()
    fmt.Printf("%s\n", pathName)
}


func (p path) ToUpper() {
    fmt.Printf("%p\n", p)
    for i, b := range p {
        if 'a' <= b && b <= 'z' {
            p[i] = b + 'A' - 'a'
        }
    }
}

