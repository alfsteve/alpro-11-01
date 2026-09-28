package main

import "fmt"

func main() {
	var posisi, posisi0, kecepatan, waktu int

	fmt.Println("Masukkan nilai dari posisi awal, kecepatan, dan waktu: ")
	fmt.Scan(&posisi0, &kecepatan, &waktu)

	posisi = posisi0 + (kecepatan * waktu)

	fmt.Println("Posisi atau jarak akhir adalah", posisi)
}