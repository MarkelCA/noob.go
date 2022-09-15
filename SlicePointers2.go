package main
import("fmt")

func main() {
    slice:= []string{"a","a"}

    func(slice []string){
        slice[0]="b";
        slice[1]="b";
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice) 

    fmt.Println("\n--------")
    func(slice []string){
        slice= append(slice, "a")
        slice[0]="b";
        slice[1]="b"; 
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice)

    fmt.Println("\n--------")
    slice2:= make([]string, 1, 3)
  
        fmt.Printf("%v\n", slice2)
    func(slice2[]string){
        slice2=slice2[1:3]
        slice2[0]="b"
        slice2[1]="b"
        fmt.Print(len(slice2))
        fmt.Print(slice2)
    }(slice2)
    fmt.Println()
    fmt.Print(len(slice2)) 
    fmt.Print(slice2)


}
