package main

import "fmt"

type person struct {
	    name string
	    age  int
	}


func main() {
    // Deklarasi dan inisialisasi map
    buah := make(map[string]int)

    // Menambahkan atau memperbarui nilai
    buah["apel"] = 5
    buah["pisang"] = 3

    // Mengakses nilai
    jumlahApel := buah["apel"]
    fmt.Println("Jumlah apel:", jumlahApel)

    // Hapus nilai
    delete(buah, "apel")

    // Cek ketersediaan kunci
    nilai, ada := buah["apel"]
    if ada {
        fmt.Println("Apel ada dalam map dan jumlahnya:", nilai)
    } else {
        fmt.Println("Apel tidak ada dalam map.")
    }

    // Iterasi map
    for kunci, nilai := range buah {
        fmt.Printf("Kunci: %s, Nilai: %d\n", kunci, nilai)
    }

    // Panjang map
    panjang := len(buah)
    fmt.Println("Panjang map buah:", panjang)
}
