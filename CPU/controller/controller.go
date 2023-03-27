package main

import (
	"demoOne/CPU/pin"
	"demoOne/CPU/utils"
	"demoOne/assembly"
	"fmt"
	"os"
)

var cpuGlobalFile *os.File

func main() {
	cpu_file, err := os.Create("./binMirco/cpu_mirco.bin")

	cpuGlobalFile = cpu_file

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

	cpu_file.Seek(0, 0)

	for addr := 0; addr < 0x10000; addr++ {
		//高8位指令
		ir := addr >> 8
		//中4位状态
		psw := (addr >> 4) & 0xf
		//获取地址的后4位确定该地址到底在哪一个周期位置
		cyc := addr & 0xf
		//如果当前指令的位置小于起始指令，那么把6条起始指令写入文件
		if cyc < len(assembly.FETCH) {
			//写入
			cpu_file.Write(uint64ToBit(uint64(assembly.FETCH[cyc]), 4))
			continue
		}

		//如果当前位在7-16位，那么开始书写指令
		addr2 := ir & (1 << 7) //获取最高位，判断是2地址指令 1xxx [aa] [bb]
		addr1 := ir & (1 << 6) //获取次高位，判断是1地址指令 01xx xx [aa]

		index := cyc - len(assembly.FETCH)

		if addr2 != 0 {
			compileAddr2(addr, ir, psw, index)
		} else if addr1 != 0 {
			compileAddr1(addr, ir, psw, index)
		} else {
			compileAddr0(addr, ir, psw, index)
		}
	}

}

// 编译指令
func compileAddr2(addr int, ir int, psw int, index int) {
	op := ir & 0xf0 //获取指令的高4位，用于确定指令的类型
	//amd := (ir) & 0xc //0xc是1100，xxxx & 0xc == xx00，所以这个操作是用来获取目的操作数
	//ams := (ir) & 0x3 //0x3是0011，xxxx & 0xc == 00xx，所以这个操作是用来获取原操作数

	newinstructions := assembly.NEWINSTRUCTIONS()

	//判断指令是否存在
	if !utils.ArrUInt32ExitNumber(&newinstructions.ADDR_2s.INSTBIT, op) { //在其中一定注册了类型，如果不存在，那么写入空指令
		//如果没有定义这个指令，那么就写入空指令
		cpuGlobalFile.Write(uint64ToBit(uint64(pin.CYC), 4))
		return
	}
	//指令存在
	if index < len(newinstructions.ADDR_2s.INST[uint32(op)]) {
		cpuGlobalFile.Write(uint64ToBit(uint64(newinstructions.ADDR_2s.INST[uint32(op)][index]), 4))
	} else {
		cpuGlobalFile.Write(uint64ToBit(uint64(pin.CYC), 4))
	}

}
func compileAddr1(addr int, ir int, psw int, index int) {

}
func compileAddr0(addr int, ir int, psw int, index int) {
	op := ir //用于确定指令的类型

	newinstructions := assembly.NEWINSTRUCTIONS()

	//判断指令是否存在
	if !utils.ArrUInt32ExitNumber(&newinstructions.ADDR_0s.INSTBIT, op) { //在其中一定注册了类型，如果不存在，那么写入空指令
		//如果没有定义这个指令，那么就写入空指令
		cpuGlobalFile.Write(uint64ToBit(uint64(pin.CYC), 4))
		return
	}
	//指令存在
	if index < len(newinstructions.ADDR_0s.INST[uint32(op)]) {
		cpuGlobalFile.Write(uint64ToBit(uint64(newinstructions.ADDR_0s.INST[uint32(op)][index]), 4))
	} else {
		cpuGlobalFile.Write(uint64ToBit(uint64(pin.CYC), 4))
	}
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
			if i < len(assembly.FETCH) {
				cpu_file.Write(uint64ToBit(uint64(assembly.FETCH[i]), 4))
			} else {
				cpu_file.Write(uint64ToBit(uint64(pin.HLT), 4))
			}
		}
	}
}

// 64位转二进制数组
// 1byte = 8bit，当传入size=4时，那么将64位整数转换为4个8位的二进制数组（舍弃高32位）
func uint64ToBit(number uint64, size int) []byte {
	var bytes []byte = make([]byte, size)
	for i := 0; i < size; i++ {
		//采用小端，低位在前，高位在后
		bytes[i] = uint8(number >> (i * 8))
	}
	return bytes[:]
}
