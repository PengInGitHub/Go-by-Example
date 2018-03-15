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

    //run the function cocurrently in a goroutine by using go func()
    //this new goroutine will execute concurrently with the calling one.
    go f("goroutine")

    //start goroutine for an anonymous function call 
    //now the two functions are running asynchronously

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    fmt.Scanln()
    fmt.Println("done")
}
