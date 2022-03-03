package stgo

import "fmt"

func Sgoto() {
    i := 0
HERE:
    fmt.Println(i)
    i++
    if i < 10 {
        goto HERE
    }
}
