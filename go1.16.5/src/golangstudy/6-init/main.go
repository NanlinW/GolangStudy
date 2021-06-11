package main

import (
	//"golangstudy/6-init/lib1" //1.导包的时候调用init
	//"golangstudy/6-init/lib2"
	_ "golangstudy/6-init/lib1" //2.匿名导包
	. "golangstudy/6-init/lib2" //3.将lib2包全部导入main包，少使用
)

func main() {
	//1.
	//lib1.Lib1Test()
	//lib2.Lib2Test()

	//2.
	//mylib2.Lib2Test()

	//3.
	Lib2Test()
}