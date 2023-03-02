package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("E:\\goTestIo\\test.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//文件打开失败
		return
	}

	defer file.Close() // 最后关闭文件，压入最后操作栈

	str := "test go io operation"

	var b = []byte(str)

	byts, err := file.Write(b)
	fmt.Println(byts)

}
