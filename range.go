package main 

import (
"fmt"

)

func PrintRange(){
    
//use range to sum the numbers in a slice. Arrays work like this too.
    nums := []int{1,2,4,32,2}  
    sum := 0 
    for _,num := range nums{
        sum += num
    } 
    fmt.Println(sum)

//use index in range
    for i,num := range nums{
        if num==32 {
           fmt.Println("index:",i) 
        }
    }

//use range on maps: range on map iterates over key/value pairs.
   kvs := map[string]string {"a":"apple", "b":"banana"}
   for k,v := range kvs{
        fmt.Printf("%s -> %s\n",k,v)
    }
}
