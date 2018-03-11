package main 

import (
"fmt"
"math"
)


//Interfaces are named collections of method signatures.


type geometry interface{
    perim() float64
    area() float64
}

type circle struct{
    radius float64
}

type rectI struct{
    width, height float64
}

func(r rectI) perim() float64{
    return 2*r.width + 2*r.height
}

func(r rectI) area() float64{
    return r.width*r.height
}

func(c circle) perim() float64{
    return 2*math.Pi*c.radius
}

func(c circle) area() float64{
    return math.Pi*c.radius*c.radius
}

func measure(g geometry){
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())    
}

func PrintInterfaces(){
    r := rectI{width:30,height:40}
    c := circle{radius:2}

    measure(r)
    measure(c)    
}
