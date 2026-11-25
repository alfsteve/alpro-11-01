package main

import "fmt"

type Player struct {
	nickname string
	score    int
}

func rataRata(T []Player, n int) float64 {
	total := 0

	for i := 0; i < n; i++ {
		total += T[i].score
	}

	return float64(total) / float64(n)
}

func idxTertinggi(T []Player, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if T[i].score > T[idx].score {
			idx = i
		}
	}

	return idx
}

func hitungDiAtas(T []Player, n int, batas float64) int {
	jumlah := 0

	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			jumlah++
		}
	}

	return jumlah
}

func main() {
	var n int
	fmt.Scan(&n)

	T := make([]Player, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	avg := rataRata(T, n)
	idx := idxTertinggi(T, n)
	jumlah := hitungDiAtas(T, n, avg)

	fmt.Printf("Rata-rata skor: %.2f\n", avg)
	fmt.Printf("Player terbaik: %s (%d)\n", T[idx].nickname, T[idx].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", jumlah)
}