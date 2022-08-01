package main

func changeData(name *string) {
	*name = "hahaha"
}
func main() {
	name := "jojo"
	changeData(&name)
	println(name)
}
