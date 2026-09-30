package main
import "fmt"

func main() {
	var y, x int

	fmt.Print("Masukkan jumlah kue: ")
	fmt.Scan(&y)

	fmt.Print("Masukkan jumlah anggota keluarga: ")
	fmt.Scan(&x)

	fmt.Print("Jumlah kue yang tersisa: ")
	fmt.Println(y % x)
	
}
