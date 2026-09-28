package main

import "fmt"

func main() {

	// Deklarasi variabel x dan y
    var x, y float64

	// Input nilai x dan y dari pengguna
	fmt.Println("Masukkan nilai x dan y:")
	fmt.Scan(&x, &y)

	// Perhitungan hasil dari persamaan
	hasil := (1.0 / (3*x*x + 10.0)) + (10.0 * y) + 7.0

	// Menampilkan hasil perhitungan
	fmt.Println(hasil)

}
