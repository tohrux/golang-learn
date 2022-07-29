//Pointer calculation
package main

import (
	"fmt"
	"unsafe"
)

//根据数组中第一项的内存地址 + 1，找到下一项的值
func main() {
	dataList := [3]int8{11, 22, 33}

	//获取数组第一个元素的地址
	var firstDataPtr *int8 = &dataList[0]

	//转换成Pointer类型
	ptr := unsafe.Pointer(firstDataPtr)

	//转换成uintptr类型，就行内存地址的计算（即: 地址加1个字节，意味着取第二个索引位置的值）
	targetAddress := uintptr(ptr) + 2

	//根据新地址，转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)

	value := (*int8)(newPtr)

	fmt.Println(*value)
}
