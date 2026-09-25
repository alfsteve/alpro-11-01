package main

	import "fmt"

	func main() {
		var celcius, reamur, fahrenheit, kelvin float32

		fmt.Print("Masukkan Celcius: ")
		fmt.Scan(&celcius)

		reamur = celcius * 4.0 / 5.0
		fahrenheit = celcius * 9.0 / 5.0 + 32.0
		kelvin = celcius + 273.15

		fmt.Print("Hasil Konversi: ")
		fmt.Println(celcius, "C =", reamur, "R", "=", fahrenheit, "F", "=", kelvin, "K")
		
} 