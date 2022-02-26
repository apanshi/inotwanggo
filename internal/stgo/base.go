package stgo

import "fmt"

func SbaseList() {
	var list2 = [5]string{"foo", "bar", "baz", "qux", "quux"}

	for _, v := range list2 {
		fmt.Println(v)
	}
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
