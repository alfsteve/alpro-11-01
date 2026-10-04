package main
import "fmt"

func main() {
	var n, a, b int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	fmt.Println(n % a == 0 && n % b == 0)
}
