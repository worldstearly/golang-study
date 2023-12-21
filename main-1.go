package main

import "fmt"

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

}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}
