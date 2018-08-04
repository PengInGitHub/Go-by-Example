package main

import (
	"fmt"
)

func fg() {
	//anonymous func
	defer func() {
		//recover must be in a defer func, in this way it would be called after panic
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		//panic if i > 3
		panic(fmt.Sprintf("%v", i))
	}
	//otherwisecall himself by g(i+1)
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	//recursive function
	g(i + 1)
}
