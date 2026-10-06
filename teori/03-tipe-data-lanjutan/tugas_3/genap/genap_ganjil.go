package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	fmt.Println(n%2 == 0)

}