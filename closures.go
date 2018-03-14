package main 

import (
"fmt"

)
//Go supports anonymous functions, which can form closures.
//Anonymous functions are useful when you want to define a function inline without having to name it.


//func returns another func
//This function intSeq returns another function, which we define anonymously in the body of intSeq. 

func intSeq() func() int{
    i := 0
    return func() int{
        i++
        return i
    }

}


func PrintClosures(){
    
    nextInc := intSeq()
    fmt.Println(nextInc())
    fmt.Println(nextInc())
    fmt.Println(nextInc())
}
