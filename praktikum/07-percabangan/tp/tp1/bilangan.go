package main
import "fmt"
func main(){
	var n int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&n)
	
	if n == 0 {
		fmt.Print("nol")
	} else {
		if n % 2 == 0 {
		fmt.Print("Genap, ")
	} else {
		fmt.Print("Ganjil, ")
	}

	if n > 0 {
		fmt.Print("Positif")
	} else {
		fmt.Print("Negatif")
	}
}
}	