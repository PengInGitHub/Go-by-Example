package main 

import (
"fmt"

)
//Variadic functions can be called with any number of trailing arguments.
//For example, fmt.Println is a common variadic function.

func sum(nums ...int){
    fmt.Print(nums, " ")
    total := 0
    for _,num := range nums{
        total += num 
    }
    fmt.Print(total)
}


func PrintVariadicFunctions(){
    sum(3,4,5)    
}
