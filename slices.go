package main 

import (
"fmt"

)

//slices are a key data type in Go, giving a more powerful interface to sequences than arrays

//to create an empty slice with non-zero length, use the builtin make

func PrintSlices(){
    
    s := make([]string, 3)
    fmt.Println("emp:", s)
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

//use builtin append() and copy()
    c := make([]string,len(s))
    copy(c,s)
    fmt.Println(c)
}
