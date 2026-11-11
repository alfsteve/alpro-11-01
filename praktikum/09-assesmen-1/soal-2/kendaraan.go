package main

import "fmt"

func main() {
	var x string
	m, k, t := 0, 0, 0

	for {
		fmt.Scan(&x)

		if x == "M" {
			m++
		} else if x == "K" {
			k++
		} else if x == "T" {
			t++
		} else {
			break
		}
	}

	fmt.Println("Tipe M =", m)
	fmt.Println("Tipe K =", k)
	fmt.Println("Tipe T =", t)
}