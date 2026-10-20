package main
import "fmt"
func main(){
	var angka int

	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&angka)

	if angka >= 85 && angka <= 100 {
		fmt.Println("Kamu dapat A")
	} else if angka >= 70 && angka < 85 {
		fmt.Println("Kamu dapat B")
	} else if angka >= 50 && angka < 70 {
		fmt.Println("Kamu dapat C")
	} else if angka < 50 {
		fmt.Println("Kamu dapat E")
	} else {
		fmt.Println("Nilai tidak valid.")
	}
}
