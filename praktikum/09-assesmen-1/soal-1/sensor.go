package main
import "fmt"
func main() {
	var angka, terakhir, besar, kecik, rata2, jumlah int

	fmt.Scan(&angka)
	besar, kecik, rata2, jumlah, terakhir = angka, angka, angka, 1, angka

	for {
		fmt.Scan(&angka)

		if angka == terakhir {
			break
		}

		if angka > besar {
			besar = angka
		}

		if angka < kecik {
			kecik = angka
		}

		rata2 += angka
		jumlah++
		terakhir = angka

	}

	fmt.Printf("%d %d %.2f\n", besar, kecik, float64(rata2)/float64(jumlah))
}