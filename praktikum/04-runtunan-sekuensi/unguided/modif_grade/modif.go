package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	switch {
		 case nilai >= 90 && nilai <= 100:
			grade = "A"
		case nilai >= 80 && nilai < 90:
			grade = "B"
		case nilai >= 70 &&	 nilai < 80:
			grade = "C"
		case nilai >= 60 && nilai < 70:
			grade = "D"
		case nilai >= 0 && nilai < 60:
			grade = "E"
		default:
			grade = "Nilai tidak valid"
	}
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Grade: %s\n", grade)
}