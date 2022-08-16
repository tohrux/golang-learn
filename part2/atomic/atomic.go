package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int64 = 100

func add() {
	atomic.AddInt64(&i, 1)
}

func sub() {
	atomic.AddInt64(&i, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(i)
}
