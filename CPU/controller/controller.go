package main

import (
	"demoOne/CPU/pin"
	"fmt"
	"os"
)

func main() {
	cpu_file, err := os.Create("./binMirco/cpu_mirco.bin")

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err2 := cpu_file.Close()
		if err2 != nil {
			fmt.Println(err2)
		}
	}()

	setHLT(cpu_file)

	cpu_file.Seek(0, 0) //设置光标偏离位置为0，位置为0（开始），1（当前），2（末尾）

	setFETCH(cpu_file)

}

// 给16位地址32位数据的bin文件写满终止指令
func setHLT(cpu_file *os.File) {
	bytes := uint64ToBit(uint64(pin.HLT), 4)
	for i := 0; i < 1<<16; i++ {
		cpu_file.Write(bytes[:])
	}
}

func setFETCH(cpu_file *os.File) {
	for j := 0; j < 0xfff; j++ {
		for i := 0; i < 16; i++ {
			if i < len(pin.FETCH) {
				cpu_file.Write(uint64ToBit(uint64(pin.FETCH[i]), 4))
			} else {
				cpu_file.Write(uint64ToBit(uint64(pin.HLT), 4))
			}
		}
	}
}

// 64位转二进制数组
func uint64ToBit(number uint64, size int) []byte {
	var bytes []byte = make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = uint8(number >> (i * 8))
	}
	return bytes[:]
}

//bytes := uint64ToBit(0x8000_0000, 4)
//
//fmt.Println(strconv.FormatUint(uint64(HLT), 2))
//fmt.Println(strconv.FormatUint(0x8000_0000, 2))
//
//for i := 0; i < 1<<16; i++ {
//	cpu_file.Write(bytes[:])
//}
