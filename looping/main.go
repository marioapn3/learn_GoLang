package main

import (
	"fmt"
)

func main() {
	//LOOPING
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	//LOOPING WITH BREAK
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Angka", i)
	}
	//LOOPING WITH CONTINUE
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Angka", i)
	}
	//LOOPING WITH LABEL
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
