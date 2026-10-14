package main
import "fmt"
func main() {
	var a, b, n int

	fmt.Print("Masukkan tabungan awal: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan kenaikan tabungan: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&n)

	total := 0;
	tabunganHari := a;

	for i := 1; i <= n; i++ {
		total = total + tabunganHari
		tabunganHari = tabunganHari + b
	}
	fmt.Println(total)
}