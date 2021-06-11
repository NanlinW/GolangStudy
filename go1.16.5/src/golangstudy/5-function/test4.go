package main

import "fmt"

func add1(a int, b int) int {
	fmt.Println("---add1---")
	var c int = 0
	c = a + b

	return c
}

//返回多个返回值，都是匿名的
func add2(a int, b int, c int) (int, int) {
	fmt.Println("---add2---")
	d := a + b
	e := b + c
	return d, e
}

//返回多个返回值，有形参名称的
func add3(a int, b int, c int) (r1 int, r2 int) {
	fmt.Println("---add3---")
	r1 = a + b
	r2 = b + c
	return
}

//返回多个返回值，形参类型相同
func add4(a int, b int, c int) (r1, r2 int) {
	fmt.Println("---add4---")
	fmt.Println("r1 = ", r1) //0
	fmt.Println("r2 = ", r2) //0

	r1 = a + b + 2*c
	r2 = a + b + c
	return
}

func main()  {
	a := 10
	b := 20
	c := 30
	d := add1(a, b)
	fmt.Println("d =", d)

	e, f := add2(a, b, c)
	fmt.Println("e =", e, "f =", f)

	ret1, ret2 := add3(a, b, c)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret3, ret4 := add4(a, b, c)
	fmt.Println("ret3 =", ret3, "ret4 =", ret4)
}