package main
import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Println("Hasil perbandingan:")
	fmt.Println("a > b:", a > b)
	fmt.Println("a == b:", a == b)
	fmt.Println("a < b:", a < b)
}