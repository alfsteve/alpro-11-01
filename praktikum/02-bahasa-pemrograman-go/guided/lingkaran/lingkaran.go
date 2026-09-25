package main

import "fmt"

func main() {

	//Membuat variabel
	var pi float64 = 3.14
	var r float64

	//Membaca input
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	//Membaca Output
	fmt.Print("Luas lingkaran: ", pi, " * ", r, " * ", r, " = ")
	fmt.Println(pi * r * r)
}