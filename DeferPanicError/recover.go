//recover
//recover adalah function yang bisa digunakan untuk menangkap data panic
//recover hanya bisa digunakan di dalam function defer
//jika tidak ada data panic, maka recover akan mengembalikan nilai nil
//jika ada data panic, maka recover akan mengembalikan nilai data panic tersebut
//recover sebaiknya digunakan di dalam function defer yang berada di dalam function yang di-panic

package main

import "fmt"

func endApp() {

	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Hello World")
}
