package assembly

import "demoOne/CPU/pin"

// 取指令微程序
var FETCH = []uint32{
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.IR_IN | pin.PC_INC,  //将RAM数据送入IR寄存器中，PC+1
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.DST_IN | pin.PC_INC, //将RAM数据送入DST寄存器中，PC+1
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.SRC_IN | pin.PC_INC, //将RAM数据送入SRC寄存器中，PC+1
}

// 移动指令
var MOV = 0 | pin.ADDR2

// 加指令
var ADD = (1 << pin.ADDR2_SHIFT) | pin.ADDR2

// 0地址指令
var HLT = 0x3f

type INSTRUCTIONS struct {
	ADDR_2 []uint32
}
