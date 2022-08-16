package main

import (
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 100; i++ {
		println("show")
	}
}
func show2() {
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		println("show2")
	}
}

func main() {
	go show2()
	go show()

	time.Sleep(time.Second)
}
