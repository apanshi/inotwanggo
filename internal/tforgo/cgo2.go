package tforgo

/*
#include <stdio.h>
void hello() {
	printf("C go2 test for new: hello, world\n");
}
*/
import "C"

func Tcgo2() {
	C.hello()
}
