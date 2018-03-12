package main 

import (
"fmt"
"time"
)

//Switch statements express conditionals across many branches.

//You can use commas to separate multiple expressions in the same case statement. 

func PrintSwitch(){
    i:=3
    switch i {
    case 1:
      fmt.Println("one")  
    case 3:
      fmt.Println("three")   
    }

    switch time.Now().Weekday(){
    case time.Saturday, time.Sunday:
      fmt.Println("it is weekend")   
    default:
      fmt.Println("it is weekday")     
    }

//switch without an expression is an alternate way to express if/else logic
    t := time.Now()
    switch{
        case t.Hour()<12:
          fmt.Println("it is before noon")     
        default:
          fmt.Println("it is after noon")     
    }

//A type switch compares types instead of values. You can use this to discover the type of an interface value
    whatAmI := func(i interface{}){
        switch t := i.(type) {
        case bool:
           fmt.Println("it is bool")
        case int:
           fmt.Println("it is int")  
        default:
           fmt.Printf("unknown type %T\n", t)    
        }
    }
    whatAmI(1)
    whatAmI("haha")
    whatAmI(false)
}
























