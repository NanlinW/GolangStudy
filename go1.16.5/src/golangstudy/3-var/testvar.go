/*
	声明变量的四种方式
*/
package main

import "fmt"
//声明全局变量方法一、二、三是可以的
var ga int = 100
var gb = 200
//方法四就不行
//gc := 300
func main() {
	//方法一：声明一个变量 默认值是0
	var a int
	fmt.Println("a =", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)

	//方法三 初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	//方法四：省去var关键字
	d := 3.14
	fmt.Printf("type of d = %T\n", d)

	e := 100
	fmt.Printf("tpye of e = %T\n", e)

	f := "abcd"
 	fmt.Printf("tpye of f = %T\n", f)

	//全局变量
	fmt.Println("ga =", ga, "gb =", gb)
	//fmt.Println("gc =",gc)

	//声明多个变量
	var aa, bb int = 100, 200
	fmt.Println("aa =", aa, "bb =", bb)

	var cc, dd = 3.14, "abcd"
	fmt.Println("cc =", cc, "dd =", dd)

	//多行的多变量声明
	var (
		ee int = 1000
		ff string = "ijkl"
	)
	fmt.Println("ee =", ee, "ff =", ff)

}