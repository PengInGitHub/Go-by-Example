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
}
























