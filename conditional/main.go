package main

import (
	"fmt"
)

func main() {
	var point = 8
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. nilai anda %d\n", point)
	}
}
