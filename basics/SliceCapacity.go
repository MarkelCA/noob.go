package main
import("fmt")

type slice []int
type array [4]int

func main(){
    p := slice ( make([]int, 0, 0) )

    fmt.Printf("%p\n",p)

    p.printAtts()
    p = append(p, 1)
    p.printAtts()
    // Compile error
    //p[1] = 2
    p.printAtts()

    fmt.Println("##########\nArray Tests\n##########\n-----------")

    ar := array( [4]int{1,2,3,4} )
    p2 := slice(ar[:2])
    ar.printAtts()
    p2.printAtts()

    printAction("Update index 1")
    p2[1] = 43
    ar.printAtts()
    p2.printAtts()
    fmt.Println(ar)

    printAction("Add one (fits capacity)")
    p2 = append(p2, 6)
    ar.printAtts()
    p2.printAtts()


    printAction("Add one (doesn't fit capacity)")
    p2 = append(p2, 7)
    ar.printAtts()
    p2.printAtts()


    printAction("Add another three")
    p2 = append(p2, 8, 9, 10)
    ar.printAtts()
    p2.printAtts()

    printAction("Add another one")
    p2 = append(p2, 11)
    ar.printAtts()
    p2.printAtts()

    printAction("Update slice's last index")
    p2[len(p2)-1] = 99
    ar.printAtts()
    p2.printAtts()

    printAction("Update slice's first index")
    p2[1] = 22
    ar.printAtts()
    p2.printAtts()
}

func printAction(a string) {
    fmt.Println(a)
    fmt.Println("--------")
}

func (a *array) printAtts() {
    fmt.Println("Array:")
    fmt.Printf("  value=%v\n",a)
    fmt.Printf("  pointer=%p\n",a)
}

func (s slice) printAtts() {
    fmt.Println("Slice:")
    fmt.Printf("  len=%d\n", len(s))
    fmt.Printf("  cap=%d\n", cap(s))
    fmt.Printf("  pointer=%p\n", s)
    fmt.Printf("  value=%v\n",s)
    fmt.Println("-----------")
}
