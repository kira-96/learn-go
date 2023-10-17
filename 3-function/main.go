package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum1, mul int
	sum1 = sum(os.Args[1], os.Args[2])
	fmt.Println(sum1)
	sum1, mul = calc(os.Args[1], os.Args[2])
	fmt.Println(sum1, mul)

	// 与C一样，& 取地址，* 取值
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
}

func sum(arg1 string, arg2 string) int {
	int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	return int1 + int2
}

func calc(arg1 string, arg2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	sum = int1 + int2
	mul = int1 * int2
	return
}

// 指针参数
// 将值传递给函数时，该函数中的每个更改都不会影响调用方。
// Go 是“按值传递”编程语言。
// 每次向函数传递值时，Go 都会使用该值并创建本地副本（内存中的新变量）

func updateName(name *string) {
	*name = "David"
}
