//panic
//panic function adalah function yang bisa digunakan untuk menghentikan program
//panic biasanya dipanggil ketika terjadi error pada saat program dijalankan
//panic akan menghentikan program, namun defer function tetap akan dieksekusi
//panic bisa diakses nilainya dengan recover function
//panic bisa digunakan untuk menghentikan proses panic

package main

import "fmt"

func endApp() {
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
}
