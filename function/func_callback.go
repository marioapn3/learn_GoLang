//function callback
//function callback adalah function yang bisa dijadikan parameter untuk function lain

package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data asli \t\t:", data)
	//data asli : [wick jason ethan]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	//filter ada huruf "o" : [jason]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
	//filter jumlah huruf "5" : [jason ethan]
}

func filter(data []string , callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}