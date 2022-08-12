// Package slice CURD
package main

import "fmt"

//delete
func del(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)

}

//update
func update(slice []int, index int, newItem int) []int {
	return append(append(slice[:index], newItem), slice[index+1:]...)
}

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{1, 2, 3, 4, 5}
	s2 := del(a1, 2)
	s3 := update(a2, 0, 2)
	fmt.Println(s2, s3)
}
