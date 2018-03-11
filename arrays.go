package main 

import (
"fmt"

)

//in Go, an array is a numbered sequence of elements of a soecific length

func PrintArray(){
    
    var a [5]int
    fmt.Println(a)

//set a value at an index using the array[index] = value
    a[4] = 100
    fmt.Println(a)

}
