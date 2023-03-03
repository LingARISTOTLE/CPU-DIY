package main

import "fmt"

func main() {

	fmt.Println("输入输出测试")
	var a int
	var b string
	fmt.Scanln(&a, &b) //　从终端读取字符串(输入时以空格分隔)
	//fmt.Scanf("%d %s", &a, &b) // 格式化从终端读取(输入时以空格分隔)

	fmt.Println(a)                 // 向终端打印
	fmt.Printf("a=%d, b=%s", a, b) // 格式化向终端打印

}
