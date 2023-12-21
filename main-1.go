package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")
	// 一个简单的累加的例子，break 和 continue 的用法与其他语言没有区别。
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}
	fmt.Println("sum: ", sum)
	// 对数组(arr)、切片(slice)、字典(map) 使用 for range 遍历：
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// 测试 方法调用
	i1 := 10
	i2 := 8
	result := add(i1, i2)
	result2, result3 := div(i1, i2)
	fmt.Println("result: ", result, "result2: ", result2, "result3: ", result3, " i1: ", i1, " i2: ", i2)

	// 错误处理(error handling)
	// 如果函数实现过程中，如果出现不能处理的错误，可以返回给调用者处理。
	// 比如我们调用标准库函数os.Open读取文件，os.Open
	// 有2个返回值，第一个是 *File，第二个是 error，
	//  如果调用成功，error 的值是 nil，如果调用失败，
	// 例如文件不存在，我们可以通过 error 知道具体的错误信息。
	data, err := os.Open("README.md")
	// readData := make()
	// data.Read()
	fmt.Println("data:", data)
	if err != nil {
		fmt.Println(err)
	}
	// 可以通过 errorw.New 返回自定义的错误
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}
