package main

// first of all  run cmd go mod init example
// 首选 粘贴 相关代码, 然后 执行 go work use .
// at last go run .

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Opt()) // Ahoy, world!
}
