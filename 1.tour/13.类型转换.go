package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(24))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
