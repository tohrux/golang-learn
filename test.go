// OCP Open-Closed Principle
package main

//Pet define a pet interface
type Pet interface {
	eat()
	sleep()
}

//Dog define two functions in a struct
type Dog struct {
	name string
	age  int
}

type Cat struct {
	name string
	age  int
}

// Dog implement Pet's method
func (dog Dog) eat() {
	println("eat")
}

// Cat implement Pet's method
func (cat Cat) eat() {
	println("cat eat")
}

func (cat Cat) sleep() {
	println("cat eat")
}

type Person struct{}

func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	cat := Cat{}
	per := Person{}
	per.care(cat)
}
