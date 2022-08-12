package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		println("子ing")
		time.Sleep(time.Second * 2)
		ch <- 1
	}()
	data := <-ch
	println(data)
	println("main terminated")
}

func test1(ch chan int) {
}
