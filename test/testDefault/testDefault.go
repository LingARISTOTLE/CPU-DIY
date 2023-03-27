package main

import (
	"demoOne/assembly"
	"fmt"
	"github.com/mcuadros/go-defaults"
)

func main() {
	addr_2 := new(ADDR_2)
	defaults.SetDefaults(addr_2)

	fmt.Println(addr_2.MOV)

	addr_1 := GetADDR2()
	fmt.Println(addr_1.MOV)

	//defaults.SetDefaults(&addr_2)
	//fmt.Println(addr_2.MOV)
	addr2 := assembly.NewADDR_2()
	fmt.Println(addr2.MOV)
}

type ADDR_2 struct {
	MOV uint32 `default:"pin.DST_W | pin.SRC_OUT"`
}

type ADDR_1 struct {
	TEST uint32 `default:"1"`
}

func GetADDR2() *ADDR_2 {
	ex := new(ADDR_2)
	defaults.SetDefaults(ex)
	return ex
}

type ADDR_0 struct {
}
