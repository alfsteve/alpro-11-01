package main
import "fmt"
func hitung(x, y int) int {
	return 2*x + y - 5
}

func main() {
	fmt.Println(hitung(10, 5))
	fmt.Println(hitung(20, 10))
	fmt.Println(hitung(30, 21))
}
