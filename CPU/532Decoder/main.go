package main

import (
	"fmt"
	"os"
)

/*
*
该脚本用来将5个二进制位表示的数字转换为32条电路中对应的通路
如：第一个0表示32位电路中第一个通，最后一个32表示32位电路中第31位电路通
*/
func main() {
	file, err := os.Create("./binMirco/532decoder.bin")

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err2 := file.Close()
		if err2 != nil {
			fmt.Println(err2)
		}
	}()

	for i := 0; i < 32; i++ {
		var value uint32 = 1 << i
		bytes := uint32ToBytes(value)
		file.Write(bytes[:])
		//uint32ToBytes(value,32)
	}
}

func uint32ToBytes(a uint32) [4]byte {
	var bytes [4]byte
	for i := 0; i < 4; i++ {
		bytes[i] = uint8(a >> (i * 8)) // 大端
		fmt.Println(bytes[i])
	}
	return bytes
}
