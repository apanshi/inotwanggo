package tforgo

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
    "fmt"
    "unsafe"
)

func Tcprint() {
    cstr := C.CString("Hello, world! from go using C\n")
    fmt.Println(cstr)
    C.puts(cstr)
    C.free(unsafe.Pointer(cstr))
}
