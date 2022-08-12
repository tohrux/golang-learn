package main

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

// FlyFish combination
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct{}

func (fish Fish) fly() {
	println("fly")
}
func (fish Fish) swim() {
	println("swim")
}
func main() {
	var fish FlyFish
	fish = Fish{}
	fish.swim()
	fish.fly()
}
