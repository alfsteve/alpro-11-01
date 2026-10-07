package main
import "fmt"

func main() {
	var intNumber, a int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a)
	intNumber = 5
	if a < intNumber {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}


}
