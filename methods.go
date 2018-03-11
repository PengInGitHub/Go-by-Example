package main 

import (
"fmt"

)
//Go supports methods defined on struct types.

type rect struct{
    width, height int
}

func(r rect) perim() int{
    return 2*r.width + 2*r.height
} 

func(r *rect) area() int{
    return r.width*r.height
}

func PrintMethods(){
    
    r := rect{width:10, height:5}
    fmt.Println(r.perim())
    fmt.Println(r.area())

    rp := &r
    fmt.Println(rp.perim())
    fmt.Println(rp.area())
}
