package stgo

import "fmt"

func SbaseList() {
    var list2 = [5]string{"foo", "bar", "baz", "qux", "quux"}

    for _, v := range list2 {
        fmt.Println(v)
    }
}

func SbaseSwitch() {
    var room = "cave"

    switch room {
    case "lake":
        fmt.Println("lake")
    case "cave":
        fmt.Println("cave")
        fallthrough
    case "underground":
        fmt.Println("underground")
    }

    var x = 10
    x -= 1
    x += 1

    for x > 0 {
        x-- // Golang 的设计者认为 C 语言中的两种自增自减（例如：++i ; i++）写法会引起歧义。所以强制规定 ++ -- 只能写在变量后面。
        fmt.Println(x)
    }

    var n = 10
    var sum = 0
    for j := 0; j <= n; j++ {
        sum = sum + j
    }
}

func SbaseSelect() {
    fmt.Println("SbaseSelect start")
    var c1, c2, c3 chan int
    var i1, i2 int
    select {
    case i1 = <-c1:
        fmt.Printf("received ", i1, " from c1\n")
    case c2 <- i2:
        fmt.Printf("sent ", i2, " to c2\n")
    case i3, ok := (<-c3): // same as: i3, ok := <-c3
        if ok {
            fmt.Printf("received ", i3, " from c3\n")
        } else {
            fmt.Printf("c3 is closed\n")
        }
    default:
        fmt.Printf("no communication\n")
    }
    fmt.Println("SbaseSelect end")
}

func SbaseSlice() {
    // 1.
    var list2 = [5]string{"foo", "bar", "baz", "qux", "quux"}
    var slice1 = list2[:]
    // 2.
    var slice2 = []string{"foo", "bar", "baz", "qux", "quux"}
    // 3.
    var slice3 = make([]string, 5)
    var slice4 = make([]string, 5, 10)
    var slice5 = slice2[1:3] // 由于切片 左闭右开（左含右不含），所以 长度为 2 [bar baz]， 容量为 4 [bar baz qux quux]。
    var slice6 = []string{"a1", "a2"}
    fmt.Println(slice6)
    copy(slice6, slice2)
    fmt.Println(slice6)

    fmt.Println(slice1[1:4])
    fmt.Println(slice2[1:4])
    fmt.Println(slice3[1:4])
    fmt.Println(slice4[1:4])

    fmt.Println(len(slice5), cap(slice5))
    fmt.Println(slice5[:])
}

func SbaseDefer() {
    fmt.Println("normal0")
    defer fmt.Println("defer1")
    fmt.Println("normal1")
    fmt.Println("normal2")
    defer fmt.Println("defer2")
    fmt.Println("normal3")
}
