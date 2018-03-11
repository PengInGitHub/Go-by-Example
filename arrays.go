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

//use the built-in function len()
    fmt.Println(len(a))

//declare and initialize an array in one line
    b := [5]int{1,2,3,4,5}
    fmt.Println(b)

//create two-dimension array
    var twoD [2][3]int
    for i:=0;i<2;i++{
        for j:=0;j<3;j++{
            twoD[i][j] = i+j
        } 
    }
    fmt.Println(twoD)
}
