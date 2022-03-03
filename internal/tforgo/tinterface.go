package tforgo

import "fmt"

/* interface syntax
type I interface {
    M()
}
*/

// 定义接口
type usber interface {
    start()
    end()
}

type Computer struct {
    name  string
    model string
}

// 实现接口的所有方法
func (cm Computer) start() {
    fmt.Printf("%s start\n", cm.name)
}

func (cm Computer) end() {
    fmt.Printf("%s end\n", cm.name)
}

// 使用接口的方法
func working(u usber) {
    u.start()
    u.end()
}

func Tinterface() {
    cm := Computer{
        name:  "Dell",
        model: "Fx1234",
    }
    working(cm)
}
