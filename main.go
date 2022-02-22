package main

import (
	"fmt"
	//第三方库
	"rsc.io/quote"

	"main/internal/stgo"
	"main/internal/tforgo"
)

func OtherMain() {
	fmt.Printf("Go 第一个程序\n")
	stgo.Smain()

	fmt.Println("Go 特殊模块")
	tforgo.Tmain()
}

func main() {
	fmt.Println("Go 第三方库")
	fmt.Println(quote.Hello())

	OtherMain()
}
