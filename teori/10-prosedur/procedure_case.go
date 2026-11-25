package main
import "fmt"

func add(a, b int32) int32 { // fungsi punya keyword tipe data kembali dan return
	return a + b
}

func eating() { // prosedur tidak punya keyword tipe data kembali dan return
	fmt.Println("I am hungry")
	fmt.Println("Ambil plate")
	fmt.Println("Ambil makanan")
	fmt.Println("Makan")
}

func hitungF(x, y int) {
	var hasil int
	hasil = 2*x + y - 5
	fmt.Println(hasil)
} // prosedur dengan parameter

func hitungFF(x, y int) int {
	var hasil int
	hasil = 2*x + y - 5
	return hasil
}

func main() {
	fmt.Println(hitungFF(2, 3)) // cara panggil fungsi dengan parameter
	x := 6
	y := 7
	hitungF(x, y) // cara panggil prosedur dengan variabel
	hitungF(x, 3) // cara panggil prosedur dengan campuran variabel dan nilai langsung
	hitungF(2, 3) // cara panggil prosedur dengan nilai langsung
	fmt.Println(add(2, 3)) // cara panggil fungsi
	eating() // cara panggil prosedur
}