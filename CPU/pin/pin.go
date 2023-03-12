package pin

// 各个寄存器在32位地址线的位置
var MSR uint32 = 1
var MAR uint32 = 2
var MDR uint32 = 3
var RAM uint32 = 4
var IR uint32 = 5
var DST uint32 = 6
var SRC uint32 = 7
var A uint32 = 8
var B uint32 = 9
var C uint32 = 10
var D uint32 = 11
var DI uint32 = 12
var SI uint32 = 13
var SP uint32 = 14
var BP uint32 = 15
var CS uint32 = 16
var DS uint32 = 17
var SS uint32 = 18
var ES uint32 = 19
var VEC uint32 = 20
var T1 uint32 = 21
var T2 uint32 = 22

// 输出
var MSR_OUT uint32 = MSR //存储器段寄存器MSR
var MAR_OUT uint32 = MAR //存储器地址寄存器MAR
var MDR_OUT uint32 = MDR //存储器数据寄存器MDR
var RAM_OUT uint32 = RAM //存储器控制器MC
var IR_OUT uint32 = IR   //指令寄存器IR
var DST_OUT uint32 = DST //目的操作数寄存器DST
var SRC_OUT uint32 = SRC //原操作数寄存器SRC
var A_OUT uint32 = A     //A寄存器
var B_OUT uint32 = B     //B寄存器
var C_OUT uint32 = C     //C寄存器
var D_OUT uint32 = D     //D寄存器
var DI_OUT uint32 = DI   //目的变址寄存器DI
var SI_OUT uint32 = SI   //原变址寄存器SI
var SP_OUT uint32 = SP   //栈指针SP
var BP_OUT uint32 = BP   //基址指针寄存器BP
var CS_OUT uint32 = CS   //代码段寄存器CS
var DS_OUT uint32 = DS   //数据段寄存器DS
var SS_OUT uint32 = SS   //栈段寄存器SS
var ES_OUT uint32 = ES   //附加段寄存器ES
var VEC_OUT uint32 = VEC //终端向量寄存器VEC
var T1_OUT uint32 = T1   //T1临时寄存器
var T2_OUT uint32 = T2   //T2临时寄存器

// 输出位在6-10位，右移5位
var _DST_SHIFT uint32 = 5

var MSR_IN uint32 = MSR << _DST_SHIFT
var MAR_IN uint32 = MAR << _DST_SHIFT
var MDR_IN uint32 = MDR << _DST_SHIFT
var RAM_IN uint32 = RAM << _DST_SHIFT
var IR_IN uint32 = IR << _DST_SHIFT
var DST_IN uint32 = DST << _DST_SHIFT
var SRC_IN uint32 = SRC << _DST_SHIFT
var A_IN uint32 = A << _DST_SHIFT
var B_IN uint32 = B << _DST_SHIFT
var C_IN uint32 = C << _DST_SHIFT
var D_IN uint32 = D << _DST_SHIFT
var DI_IN uint32 = DI << _DST_SHIFT
var SI_IN uint32 = SI << _DST_SHIFT
var SP_IN uint32 = SP << _DST_SHIFT
var BP_IN uint32 = BP << _DST_SHIFT
var CS_IN uint32 = CS << _DST_SHIFT
var DS_IN uint32 = DS << _DST_SHIFT
var SS_IN uint32 = SS << _DST_SHIFT
var ES_IN uint32 = ES << _DST_SHIFT
var VEC_IN uint32 = VEC << _DST_SHIFT
var T1_IN uint32 = T1 << _DST_SHIFT
var T2_IN uint32 = T2 << _DST_SHIFT

// 四位控制端
// 程序可以通过两种方式操作数据：第一种是通过所有控制位来操作，第二种是通过指令的参数来操控
// 这里的四位就是用来控制控制位读写vs指令第一个参数读写vs指令第二个参数读写
var SRC_R uint32 = 1 << 10
var SRC_W uint32 = 1 << 11
var DST_R uint32 = 1 << 12
var DST_W uint32 = 1 << 13

// 程序计数器PC
var PC_WE uint32 = 1 << 14
var PC_CS uint32 = 1 << 15
var PC_EN uint32 = 1 << 16

var PC_OUT = PC_CS                 // PC数据输出信号
var PC_IN = PC_CS | PC_WE          // PC自定义增量信号
var PC_INC = PC_CS | PC_WE | PC_EN // PC+1自增信号

// 终止指令
var HLT uint32 = 1 << 31

//指定标记
/**
二地址指令：
	1xxx [aa] [bb]
一地址指令：
	01xx xx [aa]
零地址指令：
	00xx xx xx
*/
var ADDR2 = 1 << 7  // 二地址指令标记
var ADDR1 = 1 << 6  // 一地址指令标记
var ADDR2_SHIFT = 4 //二地址指令后面有4位表示两个地址
var ADDR1_SHIFT = 2 //一地址指令后面有2位表示一个地址

var AM_INS = 0 // 立即数寻址
var AM_REG = 1 // 寄存器寻址
var AM_DIR = 2 // 寄存器直接寻址
var AM_RAM = 3 // 寄存器间接寻址
