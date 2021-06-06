package main

import "fmt"

func main() {
	//转义字符的示例
	//\t 制表符
	fmt.Println("nanlin\twang")
	fmt.Println("nanlin\twang\tnanlin\twang\tnanlin\twang\tnanlin\twang\tnanlin")
	// \n换行符
	fmt.Println("今天天气是：\n")
	fmt.Println("阴天")
	// \\一个'\'符
	fmt.Println("\\user\\abc\\aaa")
	// \"一个"符号
	fmt.Println(" a\"a\"a")
	// \r回车 从当前行的最前面开始输出
	fmt.Println("天龙八部雪山飞虎\r张飞")

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")

	fmt.Println("hellohellohellohellohellohellohellohellohellohello" +
		"hellohellohellohellohellohellohellohellohellohellohellohellohello" +
		"hellohellohellohellohellohellohellohellohellohellohello")

}
