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

//use slice
    fmt.Println(c[1:3])

//declare and initialize a variable for slice in a single line
    h := []string{"are","you","ok","?"}
    fmt.Println(h)

//compose slice into multi-dimension data structure
    twoDSlice := make([][]int,3) 
    for i:=0; i<3; i++{
        innerLen := i+1
        twoDSlice[i] = make([]int, innerLen)
        for j:=0; j<innerLen; j++{
            twoDSlice[i][j] = i+j    
        }
    }
     fmt.Println("2d", twoDSlice)
}
