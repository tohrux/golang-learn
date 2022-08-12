package main

func swap(a, b int) {
	c := a
	a = b
	b = c
}

func main() {
	v1 := "jk"
	v2 := &v1
	println(v1, v2, *v2)

	v1 = "jojo"
	println(v1, v2, *v2)

}
