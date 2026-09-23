package main

	import "fmt"

	func main() {
		var a, b, c int

		//Membaca input
		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan nilai b: ")
		fmt.Scan(&b)
		fmt.Print("Masukkan nilai c: ")
		fmt.Scan(&c)

		//Menampilkan output
		fmt.Print("Hasilnya : ")
		fmt.Println(a * 10000 + b * 5000 + c * 1000)
		
	}