package stgo

import "fmt"

func Siota() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		h = iota
		i
	)

	fmt.Printf("%d %d %d %s %s %d %d %d\n", a, b, c, d, e, f, h, i)

	const (
		g = 1 << iota
		j = 3 << iota
		k
		l
	)

	fmt.Printf("g=%d j=%d k=%d l=%d\n", g, j, k, l)
}
