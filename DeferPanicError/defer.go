//defer
//function yang bisa dijadwalkan untuk dieksekusi setelah sebuah fungsi selesai dieksekusi
//defer bisa dijadwalkan lebih dari satu
//defer dijalankan secara LIFO (Last In First Out)
//defer biasanya digunakan untuk clean up resource
//defer bisa mengakses return value sebuah function
//defer bisa mengubah return value sebuah function
//defer bisa mengubah variabel yang di return sebuah function
//defer bisa mengakses variabel di luar function

package main

import (
	"fmt"

)


func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run Application")
	result := 10/value
	fmt.Println("Result", result)
}

func main(){
	runApplication(10)
}