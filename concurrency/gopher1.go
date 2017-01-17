package main

import (
	"fmt"
	"time"
)

//START OMIT
func main() {
	go foo()
	go bar()
	fmt.Println("Hello Gopher!")
	time.Sleep(5 * time.Second)
	fmt.Println("Bye Gopher!")
}

//END OMIT
func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
