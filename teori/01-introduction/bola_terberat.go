package main

import "fmt"

func main() {

	var b1, b2, b3, b4 int

	fmt.Print("Bola berat 1: ")
	fmt.Scan(&b1)

	fmt.Print("Bola berat 2: ")
	fmt.Scan(&b2)

	fmt.Print("Bola berat 3: ")
	fmt.Scan(&b3)

	fmt.Print("Bola berat 4: ")
	fmt.Scan(&b4)

	// Menentukan kandidat bola terberat
	kandidat := b1
	nomorBola := 1

	if b2 > kandidat {
		kandidat = b2
		nomorBola = 2
	}

	if b3 > kandidat {
		kandidat = b3
		nomorBola = 3
	}

	if b4 > kandidat {
		kandidat = b4
		nomorBola = 4
	}

	fmt.Println("Bola terberat adalah bola nomor", nomorBola, "dengan berat", kandidat)

}