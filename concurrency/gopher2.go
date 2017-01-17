package main

import (
	"fmt"
	"time"
)
//MAIN OMIT
func main() {
	c := make(chan int)  //declared and initialzied a channel of type int
	go func() {
		for i := 0; i < 10; i++ {
			c <- i   // sending a int value on c chan int
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)  // receiving a int value from the channel
		}
	}()
	time.Sleep(time.Second)
}
//ENDMAIN OMIT
