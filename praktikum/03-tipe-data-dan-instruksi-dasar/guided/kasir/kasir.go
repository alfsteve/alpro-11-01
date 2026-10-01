package main
import "fmt"
func main(){

	var uang int
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&uang)

	var sepuluhRibu int = uang / 10000
	var sisa int = uang % 1

	var limaRibu int = uang / 5000
	sisa = sisa % 5000

	var seRibu int = uang / 1000
	
	fmt.Println(sepuluhRibu, limaRibu, seRibu)
}