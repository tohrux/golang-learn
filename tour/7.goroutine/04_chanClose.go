package main

import "time"

func main() {
	ch1 := make(chan int)

	go sendData(ch1)

	for {
		v, ok := <-ch1
		if !ok {
			println("all done")
			break
		}
		println("读取的数据", v)
	}
}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch1 <- i
	}
	close(ch1)
}
