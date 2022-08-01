package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// Scale *Vertex 可以改变副本的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
}
