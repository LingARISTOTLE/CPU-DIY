package main

import (
	"fmt"
	"os"
)

func main() {
	cpu_file, err := os.Create("./binMirco/cpu_mirco.bin")

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err2 := cpu_file.Close()
		if err2 != nil {
			fmt.Println(err2)
		}
	}()

}
