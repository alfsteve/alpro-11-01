package main

	import "fmt"

	func main() {
		var a int
		var sepuluh, lima, satu int
		var b int

		//Membaca input
		fmt.Print("Masukkan nilai uang: ")
		fmt.Scan(&a)

		//Menampilkan output
		sepuluh = a / 10000
		b = a % 10000
		lima = b / 5000
		b = b % 5000
		satu = b / 1000

		fmt.Println("Pecahan uang:")
		fmt.Println("10.000:", sepuluh)
		fmt.Println("5.000:", lima)
		fmt.Println("1.000:", satu)
		
	}