package main
import "fmt"
func main(){

	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)


	for i := 1; i <= b; i++ {
	
		hasil := a * i
		fmt.Println(a, "x", i, "=", hasil)

	}
}