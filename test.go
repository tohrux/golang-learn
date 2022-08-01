package main

import "fmt"

type Vertex3 struct {
	X, Y float64
}

func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex3) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex3{2, 3}
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := Vertex3{2, 3}
	(&p).scale(12)
	fmt.Println(p)
}
