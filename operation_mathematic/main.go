// calculate the sum of two integers using the built-in math package in Go
package main

import (
	"fmt"
)

func main() {
	var x int = 10

	var value = ((4+56)*2 - 10) / 2
	fmt.Println(value)

	var isEqual = (12 == x)
	fmt.Println(isEqual)

	var myTrue = true
	var myFalse = false
	var myBool = myTrue && myFalse
	fmt.Println(myBool)
}
