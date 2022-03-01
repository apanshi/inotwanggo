package tforgo

import "fmt"

func Tclosure() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()
	j *= 2
	a()
}

func Tclosure2() {
	fmt.Println("闭包测试 开始")
	res := addUpper()  // 执行 addUpper() 函数，返回一个闭包
	fmt.Println(res()) // 输出 2
	fmt.Println(res()) // 输出 3
	fmt.Println(res()) // 输出 4
	fmt.Println(res()) // 输出 5
	fmt.Println("闭包测试 结束")
}

func addUpper() func() int {
	x := 1
	return func() int {
		x++ //匿名函数中用到了 addUpper() 的变量 x，所以这是一个闭包
		return x
	}
}
