package main

import (
	"fmt"
	"main/internal/stgo"
	"main/internal/tforgo"
)

func main() {
	fmt.Printf("Go 第一个程序\n")
	stgo.Shelloworld()

	fmt.Printf("Go 并发测试\n")
	tforgo.Tparaclac()

	fmt.Printf("Go 反射测试\n")
	tforgo.Treflect()

	fmt.Printf("Go 调用 C 测试\n")
	tforgo.Tcprint()
}
