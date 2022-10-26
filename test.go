package main

import "fmt"

type Name struct {
	jojo string
}

func (this Name) Print() {
	println("p1")
}
func (this *Name) Print2() {
	println("p2")
}

type Employee struct {
	ID       int
	Name     string
	Salary   string
	Position string
}

var dilbert Employee
var ip *string

func main() {
	fmt.Printf("%x", ip)
}
