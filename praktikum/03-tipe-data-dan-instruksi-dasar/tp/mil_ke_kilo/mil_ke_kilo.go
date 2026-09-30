package main
import "fmt"
func main() {

	var mil float64

	fmt.Print("Masukkan jumlah mil: ")
	fmt.Scan(&mil)

	fmt.Printf("Hasil konversinya sebagai berikut: %.1f km", mil * 1.6)
}