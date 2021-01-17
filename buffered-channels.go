package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3 error occured -> "goroutines are asleep - deadlock!""
	fmt.Println(<-c)
	fmt.Println(<-c)
}
