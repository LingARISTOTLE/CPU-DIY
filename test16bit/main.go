package main

import (
	"fmt"
	"strconv"
)

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func main() {
	hex := "3f2de01"
	dec := Hex2Dec(hex)
	fmt.Println(dec)
}
