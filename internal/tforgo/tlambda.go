package tforgo

import (
    "fmt"
)

type Stringy func() string

func foo() string {
    return "Stringy function"
}

func takeAFunction(foo Stringy) {
    fmt.Println("takeAFunction: %v\n", foo())
}

func returnAFunction() Stringy {
    return func() string {
        fmt.Printf("Inner stringy function\n")
        return "bar" // have to return a string to be stringy
    }
}

func Tlambda() {
    takeAFunction(foo)
    var f Stringy = returnAFunction()
    f()
    var baz Stringy = func() string {
        return "anonymous stringy\n"
    }
    fmt.Printf("baz: %v\n", baz())
}

func Tlambda2() {
    nextInt_t := func() func() int {
        x := 0
        return func() int {
            x++
            return x
        }
    }

    nextInt := nextInt_t()

    fmt.Printf("tlambda2 start!\n")
    fmt.Println("nextInt:", nextInt()) // 1
    fmt.Println("nextInt:", nextInt()) // 2
    fmt.Println("nextInt:", nextInt()) // 3
    fmt.Println("nextInt:", nextInt()) // 4
    fmt.Println("nextInt:", nextInt()) // 5
    fmt.Printf("tlambda2 end!\n")
}
