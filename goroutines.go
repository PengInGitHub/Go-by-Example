package main 

import (
"fmt"

)

//A goroutine is a lightweight thread of execution.
//let programms run asynchronously

func f(from string){
    for i:=0;i<3;i++{
        fmt.Println(from,":",i)
    }
}

func PrintGoroutines(){
    //call a function in the usual way, running it synchronously.
    f("ohlala")
}
