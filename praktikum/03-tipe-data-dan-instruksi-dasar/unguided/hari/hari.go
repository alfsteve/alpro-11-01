package main
import "fmt"
func main(){

	var hari, Tahun, Minggu, Bulan int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	Tahun = hari / 360
	hari = hari % 360

	Bulan = hari / 30
	hari = hari % 30

	Minggu = hari / 7
	hari = hari % 7

	fmt.Println("Tahun =", Tahun)
	fmt.Println("Bulan =", Bulan)
	fmt.Println("Minggu =", Minggu)
	fmt.Println("Hari =", hari)

}