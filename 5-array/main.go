package main

import (
	"fmt"
)

func main() {
	// 声明数组
	var a [3]int
	a[1] = 10
	fmt.Println(a[:len(a)-1])

	// 初始化数组
	// 可使用省略号自动设置长度
	cities := [...]string{"北京", "上海", "广州", "深圳", "杭州"}
	fmt.Println("Cities:", cities)

	// 多维数组
	var a1 [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			a1[i][j] = (i + 1) * (j + 1)
		}
	}
	fmt.Println("2D array:", a1)

	// 切片
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	quarter2Ext := quarter2[:6]
	fmt.Println(quarter2Ext, len(quarter2Ext), cap(quarter2Ext))
	// 追加项
	quarter2Ext = append(quarter2Ext, "1") // 会改变原数组
	fmt.Println(quarter2Ext, len(quarter2Ext), cap(quarter2Ext))
	quarter2Ext[6] = "October" // 恢复
	fmt.Println(months)
	// 删除项，不提供函数，自己想办法
	// 创建副本，不更改原数组
	slice2 := make([]string, 3)
	copy(slice2, quarter2[:3])
	slice2[0] = "A"
	fmt.Println(slice2)
	fmt.Println(months)

	// 映射 map[T]T
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	// 访问不存在的项不会返回错误
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found.")
	}
	// 删除项
	// 删除不存在的项不会返回错误
	delete(studentsAge, "christy")
	// 循环
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
