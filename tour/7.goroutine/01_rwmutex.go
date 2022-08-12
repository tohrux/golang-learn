package main

import (
	"fmt"
	"sync"
	"time"
)

var RWMutex *sync.RWMutex
var wg2 *sync.WaitGroup

func main() {
	RWMutex = new(sync.RWMutex)
	wg2 = new(sync.WaitGroup)

	wg2.Add(2)

	//同时读取
	go readData(1)
	go readData(2)

	wg2.Wait()
	fmt.Println("main...over...")
}
func readData(i int) {
	defer wg2.Done()

	fmt.Println(i, "开始读。。。")
	RWMutex.RLock() //读操作上锁
	fmt.Println(i, "正在读取数据")
	time.Sleep(time.Second * 3)
	RWMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束")
}
