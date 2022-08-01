package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func something(arg interface{}) {
	tp, ok := arg.(Person)
	if ok {
		fmt.Println(tp.name)
	} else {
		fmt.Println("转换失败")
	}
}
func some2(arg interface{}) {
	switch t := arg.(type) {
	case Person:
		fmt.Println(t.name)
	}
}

// IBase 定义接口
type IBase interface {
	f1() int
}
type Person2 struct {
	name string
	age  int
}

func (v Person) f1() int {
	return 1
}
func (v Person2) f1() int {
	return 2
}

func some3(base IBase) {
	fmt.Println(base.f1())
}

func main() {
	something(Person{name: "asd", age: 12})
	v := Person{"hell", 12}
	some2(v)

	person := Person{"person", 1}
	person2 := Person2{"person2", 2}
	some3(person2)
	some3(person)

}
