package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("E:\\goTestIo\\bin256.bin", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			temp := i*16 + j
			file.Write([]byte(string((temp/100)*16*16 + (temp/10%10)*16 + temp%10)))
			fmt.Print(string((temp/100)*16*16+(temp/10%10)*16+temp%10) + " ")

		}
		file.Write([]byte("\n"))
		fmt.Println()
	}

}
