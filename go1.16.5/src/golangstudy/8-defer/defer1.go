package main

import "fmt"

//defer的执行顺序
func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
//顺序是C B A，入栈 先入后出
func main() {
	defer func1()
	defer func2()
	defer func3()
}
