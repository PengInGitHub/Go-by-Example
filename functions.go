package main 

import (
"fmt"

)
func plus(a ,b int) int {
    return a+b
}

func PrintFunctions(){
//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.    
    

    fmt.Println(plus(3,4))
}
