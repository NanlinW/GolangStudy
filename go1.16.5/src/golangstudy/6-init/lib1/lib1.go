package lib1

import "fmt"

func Lib1Test() { //模块中要导出的函数，必须首字母大写
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1. init() ....")
}
