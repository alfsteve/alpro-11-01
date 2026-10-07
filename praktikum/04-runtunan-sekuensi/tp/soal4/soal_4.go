package main
import "fmt"

func main() {
	var warna string
	fmt.Print("Masukkan warna: ")
	fmt.Scan(&warna)

	switch warna {
		case "Merah":
			fmt.Println("Aku marahhh >:(")
		case "Kuning":
			fmt.Println("Aku sedang khawatir :(")
		case "Hijau":
			fmt.Println("Aku merasa tenang :)")
		default:
			fmt.Println("Warna tidak dikenali mas")
	}
}