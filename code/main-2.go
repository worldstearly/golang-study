package main

import (
	"fmt"
)

func main() {
	// 结构体(struct) 和方法(methods)
	fmt.Println("结构体(struct) 和方法(methods)")
	fmt.Println("结构体类似于其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等。\r\n接下来我们定义一个结构体 Student，并为 Student 添加 name，age 字段，并实现 hello() 方法。")
	// 声明 stu 变量 是 取地址 的 Student 类型，其中 name 为 Tom
	stu := &Student{
		name: "Tom",
	}
	// 使用 Student{field: value, ...}的形式创建 Student 的实例，字段不需要每个都赋值，没有显性赋值的变量将被赋予默认值，例如 age 将被赋予默认值 0。
	fmt.Println(stu.hello("Jack"))
}

/*
*
声明一个 Student 类型的结构体 有name 和 age 字段
*/
type Student struct {
	name string
	age  int
}

// 声明方法 hello 参数接收一个 string  返回一个 string
func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}
