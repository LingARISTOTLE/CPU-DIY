package pin

// 取指令微程序
var FETCH = []uint32{
	PC_OUT | MAR_IN,           //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	RAM_OUT | IR_IN | PC_INC,  //将RAM数据送入IR寄存器中，PC+1
	PC_OUT | MAR_IN,           //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	RAM_OUT | DST_IN | PC_INC, //将RAM数据送入DST寄存器中，PC+1
	PC_OUT | MAR_IN,           //把pc中的数送到总线上，并将总线上的数送到MAR寄存器中
	RAM_OUT | SRC_IN | PC_INC, //将RAM数据送入SRC寄存器中，PC+1
}
