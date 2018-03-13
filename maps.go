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

    
}
