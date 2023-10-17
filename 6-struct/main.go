package main

import (
	"fmt"
	"strings"
)

// 结构体
// 在 Go 中，只需使用大写标识符，即可公开方法(public)
// 使用非大写的标识符将方法设为私有方法(private)。
// Go 中的封装仅在程序包之间有效。
type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string
}

type Employee struct {
	Person    // 继承？
	ManagerID int
}

type Triangle struct {
	size int
}

func (t *Triangle) perimeter() int {
	return t.size * 3
}

type coloredTriangle struct {
	Triangle
	color string
}

func (t *coloredTriangle) perimeter() int {
	return t.Triangle.perimeter() * 2
}

type stringext string

func (s stringext) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	// 结构体
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName, employee.LastName)

	t := coloredTriangle{Triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter (normal)", t.Triangle.perimeter())
	fmt.Println("Perimeter (colored)", t.perimeter())

	s := stringext("Learning Go!")
	fmt.Println(s.Upper())
}
