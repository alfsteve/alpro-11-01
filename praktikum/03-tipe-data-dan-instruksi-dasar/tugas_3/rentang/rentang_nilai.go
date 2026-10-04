package main
import "fmt"

func main() {
	var x, low, high int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan low: ")
	fmt.Scan(&low)
	fmt.Print("Masukkan bilangan high: ")
	fmt.Scan(&high)

	fmt.Println(x >= low && x <= high)
}
