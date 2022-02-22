package tforgo

import "C"

func HelloGo(value string) *C.char {
	return C.CString("Hello from gogogo " + value)
}
