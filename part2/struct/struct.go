package main

import "fmt"

type Person struct {
	name string
}

func (m Person) HelloWorld() {
	println(m.name)
	println("HELL")
}

func main() {
	var p Pet
	p = Dog{}
	p = Cat{}
	p.eat()
}

type Pet interface {
	eat()
}
type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}
