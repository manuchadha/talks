package main

import "fmt"

func main() {

	//INIT OMIT
	// A channel which allows "int" type of message(Unbuffered channel).
	var c chan int
	c = make(chan int)
	// or Buffered chan of int 
	c := make(chan int,100)
	//ENDINIT OMIT

	//SEND OMIT
	// Sends a value on the channel
	c <- 1
	//ENDSEND OMIT

	//WAIT OMIT
	// Waits to receive a value on the channel
	<- c
	//ENDWAIT OMIT

	
	//REC OMIT
	//Waits to receive a value and stores it in x
	//The "arrow" indicates the direction of data flow.
	x = <- c
	//or 
	//Waits to receive a value; ok will be false if channel is closed and empty
	x, ok = <- c
	//ENDREC OMIT

}
