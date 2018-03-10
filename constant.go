package main 

import (
"fmt"
"math"
)

const s = "constant"

func PrintConstant(){
    const n = 400000000
    
    fmt.Println(s,int(n)/4,math.Sin(n))
}
