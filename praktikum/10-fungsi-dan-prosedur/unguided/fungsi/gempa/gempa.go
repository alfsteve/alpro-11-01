package main

import "fmt"

type Gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	}

	return "Besar"
}

func rataMagnitudo(T []Gempa, n int) float64 {
	total := 0.0

	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}

	return total / float64(n)
}

func idxTerkuat(T []Gempa, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idx].magnitudo {
			idx = i
		}
	}

	return idx
}

func hitungKategori(T []Gempa, n int, k string) int {
	jumlah := 0

	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}

	return jumlah
}

func main() {
	var n int
	fmt.Scan(&n)

	T := make([]Gempa, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n",
			T[i].nama, T[i].tahun,
			T[i].magnitudo, kategori(T[i].magnitudo))
	}

	avg := rataMagnitudo(T, n)
	idx := idxTerkuat(T, n)
	jumlah := hitungKategori(T, n, "Besar")

	fmt.Printf("Rata-rata magnitudo: %.2f\n", avg)
	fmt.Printf("Gempa terkuat: %s %d\n",
		T[idx].nama, T[idx].tahun)
	fmt.Printf("%.1f\n", T[idx].magnitudo)
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", jumlah)
}