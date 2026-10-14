package main
import "fmt"
func main(){

	var n int
	var a int = 0
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		a += i

	}
	fmt.Println()
	fmt.Println("hasil :", a)
}
