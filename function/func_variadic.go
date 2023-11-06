package main

import (
	"fmt"
)

//fungsi variadic adalah fungsi yang bisa menerima banyak input parameter

func main3() {
	var avg = avgCalculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func avgCalculate(number ...int) float64 {
	total := 0
	for _, number := range number {
		total = total + number
	}
	avg := float64(total) / float64(len(number))
	return avg
}
