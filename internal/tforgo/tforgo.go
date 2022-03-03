package tforgo

import "fmt"

func Tmain() {
	fmt.Printf("Go 并发测试\n")
	Tparaclac()

	fmt.Printf("Go 反射测试\n")
	Treflect()

	fmt.Printf("Go 调用 C 测试\n")
	Tcprint()

	Tclosure()
	Tthread()
	Tcgo2()
	TjsonTest()
	Tclosure2()
	Tinterface()
	Tclass()
	Tlambda()
	Tlambda2()
}
