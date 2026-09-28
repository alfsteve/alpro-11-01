package main

import "fmt"

func main() {
	var umur int8
	var suhu float32

	//umur = 10
	//suhu = 36.3

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan suhu: ")
	fmt.Scan(&suhu)

	fmt.Println("Umur saya adalah", umur, "tahun")
	fmt.Println("Suhu tubuh saya adalah", suhu, "derajat celcius")

	fmt.Println("Alamat memori dari variabel umur adalah", &umur)
	fmt.Println("Alamat memori dari variabel suhu adalah", &suhu)
}