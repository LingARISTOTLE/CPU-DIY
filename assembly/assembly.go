package assembly

import (
	"demoOne/CPU/pin"
)

// 取指令微程序
var FETCH = []uint32{
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.IR_IN | pin.PC_INC,  //将RAM数据送入IR寄存器中，PC+1
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.DST_IN | pin.PC_INC, //将RAM数据送入DST寄存器中，PC+1
	pin.PC_OUT | pin.MAR_IN,               //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	pin.RAM_OUT | pin.SRC_IN | pin.PC_INC, //将RAM数据送入SRC寄存器中，PC+1
}

// 移动指令 1000 0000
var MOV uint32 = 0 | pin.ADDR2

// 加指令  1001 0000
var ADD uint32 = (1 << pin.ADDR2_SHIFT) | pin.ADDR2

// 0地址指令 0011 1111
var HLT uint32 = 0x3f

// 空指令
var NOP uint32 = 0

type INSTRUCTIONS struct {
	ADDR_2s *ADDR_2
	ADDR_1s *ADDR_1
	ADDR_0s *ADDR_0
}

func NEWINSTRUCTIONS() *INSTRUCTIONS {

	return &INSTRUCTIONS{
		ADDR_0s: &ADDR_0{
			INST: map[uint32][]uint32{
				NOP: {pin.CYC}, // 本条指令清零指令
				HLT: {pin.HLT},
			},
			INSTBIT: []uint32{
				NOP,
				HLT,
			},
		},
		ADDR_2s: &ADDR_2{
			INST: map[uint32][]uint32{
				MOV: []uint32{pin.DST_W | pin.SRC_OUT}, //取值指令把DST寄存器中的值移动到SRC中
			},
			INSTBIT: []uint32{
				MOV,
			},
		},
	}
}

type ADDR_2 struct {
	INST    map[uint32][]uint32 //用于存储指令的值
	INSTBIT []uint32            //用于存储指令类型
}

func NewADDR_2() *ADDR_2 {
	return &ADDR_2{
		INST: map[uint32][]uint32{
			MOV: []uint32{pin.DST_W | pin.SRC_OUT}, //取值指令把DST寄存器中的值移动到SRC中
		},
		INSTBIT: []uint32{
			MOV,
		},
	}
}

type ADDR_1 struct {
	TEST uint32
}

// 0地址指令，那么只需要一条指令即可
type ADDR_0 struct {
	INST    map[uint32][]uint32 //用于存储指令的值
	INSTBIT []uint32            //用于存储指令类型
}

func NewADDR_0() *ADDR_0 {
	return &ADDR_0{
		INST: map[uint32][]uint32{
			NOP: {pin.CYC}, // 本条指令清零指令
			HLT: {pin.HLT},
		},
		INSTBIT: []uint32{
			NOP,
			HLT,
		},
	}
}
