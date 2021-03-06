package main 

import (
"fmt"

)

//Maps are Go’s built-in associative data type
//sometimes called hashes or dicts in other languages

func PrintMaps(){

//To create an empty map, use the builtin make: make(map[key-type]val-type).    
    m := make(map[string]int)
    m["k1"] = 20
    m["k2"] = 8
    fmt.Println(m)

//The builtin len returns the number of key/value pairs when called on a map         

    fmt.Println(len(m))

//The builtin delete removes key/value pairs from a map.
    delete(m,"k2")
    fmt.Println(m)
//if key presents
    _,prs := m["k2"]
    fmt.Println("prs",prs)

//init a map in the same line
    n := map[string]int{"foo":1,"bar":2}   
    fmt.Println(n) 
}
