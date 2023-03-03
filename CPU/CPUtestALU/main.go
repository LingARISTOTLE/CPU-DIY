package main

import (
	"fmt"
	"os"
	"strconv"
)

// A寄存器的W/R读写控制端和CS片选控制端
var WE_A uint16 = 1
var CS_A uint16 = 1 << 1

var WE_B uint16 = 1 << 2
var CS_B uint16 = 1 << 3

var WE_C uint16 = 1 << 4
var CS_C uint16 = 1 << 5

var ALU_ADD uint16 = 0
var ALU_SUB uint16 = 1 << 6
var ALU_OUT uint16 = 1 << 7

var WE_M uint16 = 1 << 8
var CS_M uint16 = 1 << 9

var WE_PC uint16 = 1 << 10
var EN_PC uint16 = 1 << 11
var CS_PC uint16 = 1 << 12

var HLT uint16 = 1 << 15

var micro = [...]uint16{
	CS_M | CS_A | WE_A | WE_PC | EN_PC | CS_PC,
	CS_M | CS_B | WE_B | WE_PC | EN_PC | CS_PC,
	ALU_OUT | CS_C | WE_C,
	CS_C | WE_M | CS_M | WE_PC | EN_PC | CS_PC,
	HLT,
}

func main() {
	//newfile, err := os.Create("E:\\dqcaogao\\micro.bin")
	newfile, err := os.Create("./binMirco/micro.bin")
	if err != nil {
		fmt.Println(err)
	}

	defer newfile.Close()

	for i := 0; i < len(micro); i++ {
		str := strconv.FormatUint(uint64(micro[0]), 16)
		//采用小端，低位在前，高位在后
		var bytes = []byte{uint8(micro[i]), uint8(micro[i] >> 8)}
		newfile.Write(bytes)
		fmt.Println(bytes)
		fmt.Println(str)
	}

	//var magic []byte = []byte{0x12, 0x31, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22}

	//str := strconv.FormatUint(uint64(micro[0]), 16)
	//fmt.Println([]byte(str))
	//fmt.Println("..............................................")
	//
	//fmt.Println(micro[0])
	//fmt.Println(uint8(micro[0]))
	//fmt.Println(uint8(micro[0] >> 8))
	//fmt.Println("..............................................")
	//
	//fmt.Println(strconv.FormatUint(uint64(micro[0]), 16))
	//fmt.Println(WE_A)
	//fmt.Println(CS_A)
	//fmt.Println(WE_B)
	//fmt.Println(CS_B)
	//output := strconv.FormatInt(int64(HLT), 16)
	//
	//fmt.Println("The hexadecimal conversion of", HLT, "is", output)
}

func int10BitTo16Bit(a int16) {

}
