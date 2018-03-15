package main 

import (
"fmt"

)

//Channels are the pipes that connect concurrent goroutines.
//You can send values into channels from one goroutine and receive those values into another goroutine.

//create a new channel with make(chan val-type)
//channels are typed by the value they convey

//send a value into channel: channel <- value 

func PrintChannels(){
    messages := make(chan string)
    //send a message "ping" to the channel created above messages via a goroutine
    go func() {messages <- "ping"}()
    
    fmt.Println()
}
