package main
import "fmt"
func main(){

	var c, reamur float32
	fmt.Print("Masukkan angka Celcius: ")
	fmt.Scan(&c)

	reamur = c * 4.0 / 5.0

	fmt.Println("Reamur =", reamur)
}