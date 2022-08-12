package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10
var wg sync.WaitGroup
var mutex sync.Mutex

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	mutex.Lock()
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "sale", ticket)
			ticket--
		} else {
			fmt.Println(name, "sale out")
			break
		}
	}
	mutex.Unlock()
}
func main() {
	wg.Add(3)
	go saleTickets("A")
	go saleTickets("B")
	go saleTickets("C")
	wg.Wait()
	println("done")
}
