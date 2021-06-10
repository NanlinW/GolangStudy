package main

import "fmt"

//iota只能配合const一起使用
const (
	b = 10 * iota
	c
	d
)
const (
	e, f = iota + 1, iota + 2
	g, h
	i, j = iota * 2, iota * 3
	k, l
)
func main()  {
	const a int = 1000
	fmt.Println("a =", a)

	//a = 100 报错
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	fmt.Println("e =", e, "f =", f)
	fmt.Println("g =", g, "h =", h)
	fmt.Println("i =", i, "j =", j)
	fmt.Println("k =", k, "l =", l)

}
