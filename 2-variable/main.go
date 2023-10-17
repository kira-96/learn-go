package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	var helloworld string = "hello world"
	fmt.Println(helloworld)
	// 初始化变量
	var (
		age    int = 32
		length int = 198
	)
	// 新变量的赋值，使用 :=
	firstName, lastName := "John", "Doe"
	fmt.Println(firstName, lastName, age, length)

	// 常量
	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	// 数据类型
	// int8 int16 int32 int64
	// uint8 uint16 uint32 uint64
	// int 32位系统为32 64位系统为64
	// 整数，默认为 0
	fmt.Println(unsafe.Sizeof(age))

	// rune -> int32
	var g rune = 'G'
	fmt.Println(g)

	// float32 float64
	// 浮点数，默认为 +0.000
	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	// bool，默认为 false
	var featureFlag bool = true
	fmt.Println(featureFlag)

	// string，默认为空值
	var number string = "-42"
	i, _ := strconv.Atoi(number)
	number = strconv.Itoa(i)
	fmt.Println(i, number)

	// 原始字符串 ``
	// 等同于Python的 R'' / r''
	fmt.Println(`Hello \nWorld!`)
}
