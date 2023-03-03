package main

// 各个寄存器在32位地址线的位置
var MSR int = 1
var MAR int = 2
var MDR int = 3
var RAM int = 4
var IR int = 5
var DST int = 6
var SRC int = 7
var A int = 8
var B int = 9
var C int = 10
var D int = 11
var DI int = 12
var SI int = 13
var SP int = 14
var BP int = 15
var CS int = 16
var DS int = 17
var SS int = 18
var ES int = 19
var VEC int = 20
var T1 int = 21
var T2 int = 22

// 输出
var MSR_OUT int = MSR //存储器段寄存器MSR
var MAR_OUT int = MAR //存储器地址寄存器MAR
var MDR_OUT int = MDR //存储器数据寄存器MDR
var RAM_OUT int = RAM //存储器控制器MC
var IR_OUT int = IR   //指令寄存器IR
var DST_OUT int = DST //目的操作数寄存器DST
var SRC_OUT int = SRC //原操作数寄存器SRC
var A_OUT int = A     //A寄存器
var B_OUT int = B     //B寄存器
var C_OUT int = C     //C寄存器
var D_OUT int = D     //D寄存器
var DI_OUT int = DI   //目的变址寄存器DI
var SI_OUT int = SI   //原变址寄存器SI
var SP_OUT int = SP   //栈指针SP
var BP_OUT int = BP   //基址指针寄存器BP
var CS_OUT int = CS   //代码段寄存器CS
var DS_OUT int = DS   //数据段寄存器DS
var SS_OUT int = SS   //栈段寄存器SS
var ES_OUT int = ES   //附加段寄存器ES
var VEC_OUT int = VEC //终端向量寄存器VEC
var T1_OUT int = T1   //T1临时寄存器
var T2_OUT int = T2   //T2临时寄存器

// 输出位在6-10位，右移5位
var _DST_SHIFT int = 5

var MSR_IN int = MSR << _DST_SHIFT
var MAR_IN int = MAR << _DST_SHIFT
var MDR_IN int = MDR << _DST_SHIFT
var RAM_IN int = RAM << _DST_SHIFT
var IR_IN int = IR << _DST_SHIFT
var DST_IN int = DST << _DST_SHIFT
var SRC_IN int = SRC << _DST_SHIFT
var A_IN int = A << _DST_SHIFT
var B_IN int = B << _DST_SHIFT
var C_IN int = C << _DST_SHIFT
var D_IN int = D << _DST_SHIFT
var DI_IN int = DI << _DST_SHIFT
var SI_IN int = SI << _DST_SHIFT
var SP_IN int = SP << _DST_SHIFT
var BP_IN int = BP << _DST_SHIFT
var CS_IN int = CS << _DST_SHIFT
var DS_IN int = DS << _DST_SHIFT
var SS_IN int = SS << _DST_SHIFT
var ES_IN int = ES << _DST_SHIFT
var VEC_IN int = VEC << _DST_SHIFT
var T1_IN int = T1 << _DST_SHIFT
var T2_IN int = T2 << _DST_SHIFT

// 四位控制端
var SRC_R int = 1 << 10
var SRC_W int = 1 << 11
var DST_R int = 1 << 12
var DST_W int = 1 << 13

// 程序计数器PC
var PC_WE int = 1 << 14
var PC_CS int = 1 << 15
var PC_EN int = 1 << 16

var PC_OUT = PC_CS
var PC_IN = PC_CS | PC_WE
var PC_INC = PC_CS | PC_WE | PC_EN

func main() {

}
