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

}
