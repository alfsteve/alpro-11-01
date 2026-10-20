package main
import "fmt"
func main(){

	var n int
	Faktor := 0
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n % i == 0{
			fmt.Print(i, " ")
			Faktor++
		}
	}

	fmt.Println()

	if Faktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}