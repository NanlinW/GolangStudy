package main

import "fmt"

//defer和return的顺序
func deferfunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnfunc() int {
	fmt.Println("return func called...")
	return 0
}

func deferAndmain() int {
	defer deferfunc()
	return returnfunc()
}
 //先return 后defer
 //defer是函数的全部生命周期结束后再出栈
func main()  {
	deferAndmain()
}