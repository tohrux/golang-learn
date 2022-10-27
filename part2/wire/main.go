package main

import "fmt"

func main() {
	conf := NewConfig()
	db := NewDB(conf)
	result := db.Find()
	fmt.Println(result)
}
