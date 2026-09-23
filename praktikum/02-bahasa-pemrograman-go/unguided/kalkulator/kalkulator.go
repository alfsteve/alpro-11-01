package main

	import "fmt"
	
	func main() {
		var a, b int

		//Membaca input
		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan nilai b: ")
		fmt.Scan(&b)

		//Menampilkan output
		fmt.Print("Hasil Penjumlahan: ")
		fmt.Println(a + b)
		fmt.Print("Hasil Pengurangan: ")
		fmt.Println(a - b)
		fmt.Print("Hasil Perkalian: ")
		fmt.Println(a * b)
		fmt.Print("Hasil Pembagian: ")
		fmt.Println(a / b)
		fmt.Print("Hasil Sisa Hasil Bagi: ")
		fmt.Println(a % b)

}