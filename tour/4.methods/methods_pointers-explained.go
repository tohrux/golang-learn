package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
