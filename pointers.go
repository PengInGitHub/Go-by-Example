package main 

import (
"fmt"

)

func zeroval(ival int){
    ival = 0
}


func zeroptr(iptr *int){
    *iptr = 0
}

func PrintPointers(){
    i := 1
    fmt.Println("initial",i)
    zeroval(i)
    fmt.Println("zeroval:",i)

//The &i syntax gives the memory address of i, i.e. a pointer to i.
    zeroptr(&i)
    fmt.Println("zeroptr:",i)
    fmt.Println("zeroptr:",&i)
}
