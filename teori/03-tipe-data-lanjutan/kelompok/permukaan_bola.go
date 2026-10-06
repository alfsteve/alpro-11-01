package main
import "fmt"

func main() {

	const pi = 22.0 / 7.0
	var r float64
	fmt.Println("Masukkan jari-jari bola:")
	fmt.Scan(&r)
	fmt.Println(4 * pi * (r * r))
	}