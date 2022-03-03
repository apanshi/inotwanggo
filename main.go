package main

import (
    "fmt"
    //第三方库
    "rsc.io/quote"

    "main/internal/stgo"
    "main/internal/tforgo"
)

func init() {
    fmt.Println("Go 初始化 int 1: ", a)
}

func init() {
    fmt.Println("Go 初始化 int 2: ", b)
}

var a = 10

const b = 100

func OtherMain() {
    fmt.Printf("Go 第一个程序\n")
    stgo.Smain()

    fmt.Println("Go 特殊模块")
    tforgo.Tmain()
}

func main() {
    defer fmt.Println("Go 第三方库, defer: ", quote.Hello())
    fmt.Println(quote.Hello())

    OtherMain()

    fmt.Println("Go 初始化 main(): ", b)
}
