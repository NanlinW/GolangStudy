package main

import "fmt"

func swap1(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a = 10
	var b = 20

	swap1(a, b)
	fmt.Println("a =", a, "b =", b)

	swap2(&a, &b)
	fmt.Println("a =", a, "b =", b)

	var p *int
	p = &a
	fmt.Println(*p)
	fmt.Println(a)

}


