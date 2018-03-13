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
//If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
    nums := []int{12,22,56,7}
    sum(nums...)    
}
