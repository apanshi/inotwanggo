package tforgo

/*
#include "hello.h"

#include <stdio.h>
#include <stdlib.h>
void hello() {
    printf("C go2 test for new: hello, world\n");
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Tcgo2() {
	C.hello()
	ret := C.SayHelloCxx()
	fmt.Println(ret)

	cs_hello := C.CString("Freom C’s hello")
	C.SayHelloC(cs_hello)
	C.free(unsafe.Pointer(cs_hello))
}
