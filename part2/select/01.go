package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		defer close(chanStr)
		defer close(chanInt)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v", r)
		default:
			println("sleep")
		}
		time.Sleep(time.Second)

	}
}
