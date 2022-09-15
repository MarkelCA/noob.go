// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"fmt"
)

type path []byte

// This works
func (p *path) TruncateAtFinalSlashPointer() {
	i := bytes.LastIndex(*p, []byte("/"))

    fmt.Printf("%p\n", *p)
    fmt.Printf("%p\n", p)
    fmt.Println("---- Pointer ----")
	if i >= 0 {
        fmt.Printf("%p\n", *p)
        fmt.Printf("%p\n", (*p)[0:i])
		*p = (*p)[0:i]
	}
}

// This doesn't work
// Probably because creating the slice returns a new pointer
func (p path) TruncateAtFinalSlashNotPointer() {
	i := bytes.LastIndex(p, []byte("/"))
    fmt.Println("---- No Pointer ----")
	if i >= 0 {
        fmt.Printf("%p\n", p)
        fmt.Printf("%p\n", p[0:i])
		p = p[0:i]
	}
}

// This works 
// That's because of slices value are its head's pointer, so it's like passing its reference
func (p path) ToUpper() {
    for i, b := range p {
        if 'a' <= b && b <= 'z' {
            p[i] = b + 'A' - 'a'
        }
    }
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.

	pathName.TruncateAtFinalSlashPointer()
    fmt.Printf("%s\n", pathName)
    //fmt.Printf("%s, %p \n", pathName, pathName)

	pathName2:= path("/usr/bin/tso") // Conversion from string to path.
	pathName2.TruncateAtFinalSlashNotPointer()
    fmt.Printf("%s\n", pathName2)

    fmt.Println("----- ToUpper -----")
    pathName = path("/usr/bin/tso")
    pathName.ToUpper()
    fmt.Printf("%s\n", pathName)
}
