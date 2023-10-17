package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

// Go语言中除了map、slice、chan外，其他类型在函数参数中都是值传递
// interface 属于是传递指针
type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

// 没有用于实现接口的关键字。
// 当 Go 中的接口具有接口所需的所有方法时，则满足按类型的隐式实现。
func (t *Square) Area() float64 {
	return t.size * t.size
}

func (t *Square) Perimeter() float64 {
	return t.size * 4
}

func (t *Square) String() string {
	return fmt.Sprintf("%T: size %f", t, t.size)
}

type Circle struct {
	radius float64
}

func (t *Circle) Area() float64 {
	return math.Pi * t.radius * t.radius
}

func (t *Circle) Perimeter() float64 {
	return 2 * math.Pi * t.radius
}

func (t *Circle) String() string {
	return fmt.Sprintf("%T: size %f", t, t.radius)
}

func printInformation(s Shape) {
	fmt.Printf("%s\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

type customWriter struct{}

type GithubResponse []struct {
	FullName string `json:"full_name"`
}

func (t customWriter) Write(p []byte) (n int, err error) {
	var res GithubResponse
	e := json.Unmarshal(p, &res)
	if e != nil {
		fmt.Println(e)
		return 0, e
	}
	for _, v := range res {
		fmt.Println(v.FullName)
	}

	return len(p), nil
}

func main() {
	s := Square{3}
	printInformation(&s)

	c := Circle{4}
	printInformation(&c)

	// http GET
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
