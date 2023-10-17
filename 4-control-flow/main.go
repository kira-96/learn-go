package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func main() {
	if num := give_me_a_number(); num < 0 {
		fmt.Println(num, "is negative.")
	} else {
		fmt.Println(num, "is posotive.")
	}

	// switch case 语句
	// case 后可以跟多个值
	// 可以使用变量、比较大小！
	// 不用写 break!
	//
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	// 省略switch条件
	// 这种用法类似于 if ... else ...
	email := regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	phone := regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is a email.")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number.")
	default:
		fmt.Println(contact, "is not recognized.")
	}

	// 使用 fallthrough 进入下一个 case 语句
	// 作用类似于C中不写break的case语句，但不会进行连续判断
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d < 50 \n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d > 100 \n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d < 200 \n", num)
	default:
	}

	// for 循环和其他语言类似，但不需要 ()
	num := 0
	for i := 1; i <= 100; i++ {
		num += i
	}
	fmt.Println("sum of 1..100 is", num)
	// go 没有 while 语句
	// 但可以改用 for 循环，并利用 Go 使预处理语句和后处理语句可选这一事实。
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for num != 5 {
		num = r.Intn(15)
		fmt.Println(num)
	}
	// 无限循环 for break continue
	for {
		fmt.Println("Waiting inside the loop...")
		if num = r.Intn(15); num == 5 {
			fmt.Println("Loop finish")
			break
		} else if num < 10 {
			continue
		} else {
			fmt.Println(num)
		}
	}

	// 在 Go 中，defer 语句会推迟函数（包括任何参数）的运行，
	// 直到 {包含 defer 语句的函数} 完成。
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	// Go 提供内置 recover() 函数，让你可以在程序崩溃之后重新获得控制权。
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	// defer 函数的一个典型用例是在使用完文件后将其关闭。
	// 内置 panic() 函数可以停止 Go 程序中的正常控制流。
	// panic 引发异常，发、但defer 语句会继续执行。
	write_file()
}

func give_me_a_number() int {
	return -1
}

func write_file() {
	file, error := os.Create("learnGo.txt")
	if error != nil {
		panic("Error: Could not create file.")
	}

	defer file.Close()

	if _, error = io.WriteString(file, "Learning Go!"); error != nil {
		panic("Error: Could not write to file.")
	}

	file.Sync()
	fmt.Println("Write file OK!")
	panic("手动引发异常.")
}
