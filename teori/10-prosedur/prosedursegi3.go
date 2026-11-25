package main

import "fmt"

func hitungLuas(tinggi, alas int, luas *float64) {
	*luas = float64(alas) * float64(tinggi) / 2
}

func main() {
	var N, alas, tinggi int
	var luas float64

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&alas, &tinggi)
		hitungLuas(alas, tinggi, &luas)
		fmt.Println(luas)
	}
}
