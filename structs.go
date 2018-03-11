package main 

import (
"fmt"

)

//Go’s structs are typed collections of fields.
//They’re useful for grouping data together to form records.

type person struct{
    name string
    age  int
}

func PrintStructs(){
    
    Bob := person{"Bob", 20}
    Alice := person{name : "Alice",
                    age:    30}

    fmt.Println(Alice)
    fmt.Println(Bob)
    fmt.Println(person{name : "Lily"})
//use dots with struct pointers - the pointers are automatically dereferenced.
    AlicePointer := &Alice
    AlicePointer.age = 23
    fmt.Println(Alice)
    fmt.Println(AlicePointer.age)


}
