package main
import "fmt"

func main(){
	var p, q int
	fmt.Print("Masukkan bilangan p: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan bilangan q: ")
	fmt.Scan(&q)

	fmt.Println(p % 2 == 0 || q % 2 == 0)
	fmt.Println(p % 2 != 0 && q % 2 != 0)
	fmt.Println(!(p == q))
}