package main

	import "fmt"

	func main() {
		var nama string
		var skor_Emteka, skor_BahasaInggris int

		//Membaca input
		fmt.Println("Masukkan nama:")
		fmt.Scan(&nama)

		fmt.Println("Masukkan skor Emteka:")
		fmt.Scan(&skor_Emteka)

		fmt.Println("Masukkan skor Bahasa Inggris:")
		fmt.Scan(&skor_BahasaInggris)
		
		//Menghitung total dan rata-rata (Pembagian bilangan bulat)

		total := skor_Emteka + skor_BahasaInggris
		rata_rata := total / 2

		//Menampilkan output

		fmt.Println(nama)
		fmt.Println(total)
		fmt.Println(rata_rata)

		//Menampilkan nama gitu

		fmt.Println("Nama :", nama)
		fmt.Println("Total Skor :", total)
		fmt.Println("Rata-rata Skor :", rata_rata)
}