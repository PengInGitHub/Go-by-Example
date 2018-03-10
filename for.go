package main 

import (
"fmt"

)

func PrintFor(){
    i := 1
    for i<3 {
        fmt.Println(i)   
        i++     
    }
    for j:=3;j<=10;j++{
        fmt.Println(j)
    }
    for {
        fmt.Println("Oh Yeah!")
        break
    }

    for n:=0;n<=5;n++{
	if n%2 ==0 {
            continue
	}
        fmt.Println(n)
    }

}
