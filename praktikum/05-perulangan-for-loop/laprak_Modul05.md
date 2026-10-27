# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

## Dasar Teori

### A. Pengertian Perulangan (Looping)
Perulangan atau looping merupakan struktur kontrol dalam algoritma pemrograman yang digunakan untuk menjalankan suatu blok kode secara berulang selama kondisi tertentu terpenuhi. Perulangan membantu menyederhanakan penulisan program sehingga instruksi yang sama tidak perlu ditulis berulang kali. Dalam bahasa pemrograman Go, perulangan menggunakan kata kunci for yang dapat diterapkan dalam beberapa bentuk sesuai kebutuhan program.

### B. Jenis-Jenis Perulangan dalam Go
Bahasa pemrograman Go memiliki beberapa bentuk perulangan, yaitu for dengan tiga komponen, for dengan kondisi, for tanpa kondisi, dan for-range. Perulangan dengan tiga komponen terdiri atas inisialisasi, kondisi, dan pembaruan nilai variabel. Perulangan dengan kondisi menjalankan instruksi selama kondisi bernilai benar, sedangkan for tanpa kondisi dapat menghasilkan perulangan tanpa henti apabila tidak diberikan mekanisme penghentian. Sementara itu, for-range digunakan untuk mengakses elemen dalam kumpulan data, seperti array, slice, string, dan map.

### C. Penerapan Perulangan For
Perulangan for dapat digunakan untuk menyelesaikan berbagai permasalahan, seperti menampilkan deret bilangan, menghitung jumlah bilangan dari satu sampai nilai tertentu, serta menjalankan instruksi berulang berdasarkan kondisi yang ditentukan. Dalam penerapannya, variabel penghitung (counter) digunakan untuk mengatur jumlah pengulangan. Kondisi perulangan menentukan kapan proses dilanjutkan atau dihentikan, sedangkan pembaruan nilai variabel dilakukan agar perulangan dapat mencapai kondisi berhenti.

### D. Perulangan Bersarang (Nested For-Loop)
Perulangan bersarang atau nested for-loop merupakan struktur perulangan yang menempatkan satu perulangan di dalam perulangan lainnya. Struktur ini digunakan untuk menyelesaikan permasalahan yang membutuhkan pengulangan bertingkat, seperti menampilkan pola karakter, matriks, dan tabel. Pada setiap pengulangan luar, perulangan dalam akan dijalankan sesuai kondisi yang telah ditentukan.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. tabungan.go

```go
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
```
#### Deskripsi
Program ini digunakan untuk menghitung total tabungan yang terkumpul selama beberapa hari. Program menerima tiga masukkan, yaitu tabungan awal, kenaikan tabungan setiap hari, dan jumlah hari. Variabel a menyimpan tabungan hari pertama, b menyimpan kenaikan tabungan per hari, dan n menyimpan jumlah hari yang akan dihitung.

Variabel total digunakan untuk menyimpan akumulasi seluruh tabungan, sedangkan tabunganHari menyimpan jumlah tabungan pada hari yang sedang dihitung. Dengan perulangan for, program menjumlahkan tabungan setiap hari ke dalam total, kemudian menambahkan kenaikan tabungan sebesar b untuk hari berikutnya. Setelah seluruh hari selesai dihitung, program menampilkan total tabungan yang terkumpul.

### 2. terbesar.go

```go
package main
import "fmt"
func main() {
	var n, max, angka int
	fmt.Scan(&n)

	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}

}

fmt.Println(max)
}

```
#### Deskripsi
Program ini digunakan untuk mencari nilai terbesar dari sejumlah bilangan bulat yang dimasukkan oleh pengguna. Program menerima masukan berupa jumlah bilangan n, kemudian membaca bilangan pertama dan menyimpannya ke dalam variabel max sebagai nilai terbesar sementara. 

Selanjutnya, program menggunakan perulangan for untuk membaca bilangan kedua hingga bilangan ke-n. Setiap bilangan yang dibaca disimpan dalam variabel angka, lalu dibandingkan dengan nilai max menggunakan percabangan if. Jika angka lebih besar daripada max, nilai max diperbarui dengan bilangan tersebut. Proses ini terus dilakukan sampai seluruh bilangan selesai diperiksa.

Setelah perulangan berakhir, program menampilkan nilai max sebagai bilangan terbesar dari seluruh masukan.

### 3. bioskop.go

```go
package main
import "fmt"
func main() {
	var t int
	fmt.Print("Masukkan t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int
		
		fmt.Printf("Input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0;

		if jenisFilm == "Reguler" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 40000
			} else {
				hargaSatuan = 50000
			}

		} else if jenisFilm == "3D"{
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 60000
			} else {
				hargaSatuan = 70000
			}
		} else if jenisFilm == "IMAX"{
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 80000
			} else {
				hargaSatuan = 90000
			}
		} 

	

	totalHarga := hargaSatuan * jumlahTiket

	if jumlahTiket >= 5 {
		diskon := (totalHarga * 10) / 100
		totalHarga = totalHarga - diskon

	}

	fmt.Printf("output total hargan transaksi ke-%d: %d\n", i, totalHarga)	

	totalPendapatan = totalPendapatan + totalHarga
	}

	fmt.Println("Output total pendapatan: ", totalPendapatan)
}

```
#### Deskripsi
Program ini digunakan untuk menghitung harga tiket bioskop berdasarkan jenis film, hari transaksi, dan jumlah tiket yang dibeli. Program pertama-tama menerima masukan berupa jumlah transaksi. Selanjutnya, program melakukan perulangan sebanyak jumlah transaksi yang dimasukkan.

Pada setiap transaksi, pengguna memasukkan jenis film, hari pembelian, dan jumlah tiket. Program menggunakan percabangan if-else untuk menentukan harga satuan tiket berdasarkan jenis film dan hari pembelian. Tiket Reguler memiliki harga Rp40.000 pada hari Senin sampai Kamis dan Rp50.000 pada hari lainnya. Tiket 3D memiliki harga Rp60.000 atau Rp70.000, sedangkan tiket IMAX memiliki harga Rp80.000 atau Rp90.000 berdasarkan kategori hari yang sama.

Setelah harga satuan ditentukan, program menghitung total harga dengan mengalikan harga tiket dan jumlah tiket. Jika pembelian mencapai lima tiket atau lebih, program memberikan diskon sebesar 10% dari total harga. Hasil transaksi ditampilkan, kemudian ditambahkan ke variabel totalPendapatan. Setelah seluruh transaksi selesai diproses, program menampilkan total pendapatan dari semua transaksi.  

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. prima.go

```go
package main
import "fmt"
func main(){

	var n int
	Faktor := 0
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n % i == 0{
			fmt.Print(i, " ")
			Faktor++
		}
	}

	fmt.Println()

	if Faktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/prima/output.png)


#### Deskripsi
Program ini digunakan untuk untuk menentukan faktor-faktor suatu bilangan bulat positif dan memeriksa apakah bilangan tersebut merupakan bilangan prima. Program menerima sebuah bilangan n sebagai masukan, kemudian menggunakan perulangan for untuk memeriksa setiap bilangan dari 1 sampai n.

Pada setiap iterasi, program menggunakan operator modulus (%) untuk memeriksa sisa pembagian n dengan bilangan yang sedang diuji. Jika sisanya nol, bilangan tersebut merupakan faktor dari n. Program kemudian menampilkan faktor tersebut dan menambah variabel penghitung jumlah faktor.

Setelah seluruh bilangan diperiksa, program menggunakan percabangan if-else untuk menentukan apakah jumlah faktor tepat dua. Jika jumlahnya dua, bilangan dinyatakan prima. Jika jumlah faktornya bukan dua, bilangan dinyatakan bukan prima. Dengan demikian, program dapat menampilkan daftar faktor sekaligus status bilangan yang diperiksa.

### 2. ecct.go

```go


package main

import "fmt"

func main() {
    var mahasiswa int
    fmt.Scan(&mahasiswa)

    sertif := 0

    for i := 1; i <= mahasiswa; i++ {
        total := 0

        for j := 1; j <= 6; j++ {
            var nilai int
            fmt.Scan(&nilai)
            total += nilai
        }

        ratarata := float64(total) / 6.0

        if ratarata >= 3.33 {
            fmt.Println("Ya")
            sertif++
        } else {
            fmt.Println("Tidak")
        }
    }

    fmt.Println(sertif)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/ecct/output.png)

#### Deskripsi
Program ini digunakan untuk untuk menentukan mahasiswa yang berhak mendapatkan sertifikat ECCT berdasarkan rata-rata enam nilai mata kuliah. Program menerima jumlah mahasiswa sebagai masukan, kemudian memproses nilai setiap mahasiswa menggunakan perulangan.

Di dalam setiap iterasi mahasiswa, program menginisialisasi variabel total dengan nilai nol. Perulangan bersarang digunakan untuk membaca enam nilai dan menjumlahkannya. Setelah semua nilai dibaca, program menghitung rata-rata dengan membagi total nilai dengan enam.

Program kemudian menggunakan percabangan if-else untuk memeriksa apakah rata-rata mahasiswa mencapai minimal 3,33. Jika memenuhi syarat, program menampilkan Ya dan menambah variabel sertif. Jika tidak memenuhi syarat, program menampilkan Tidak. Setelah seluruh mahasiswa selesai diproses, program menampilkan jumlah mahasiswa yang memperoleh sertifikat.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan program yang telah dibuat, dapat disimpulkan bahwa penyelesaian masalah dalam praktikum ini dilakukan dengan menentukan input, proses, dan output terlebih dahulu. Penggunaan variabel, tipe data, perulangan, serta struktur percabangan membantu program mengolah data, melakukan perhitungan, dan menghasilkan output sesuai dengan kondisi yang diberikan. Program-program tersebut menerapkan konsep dasar analisis masalah dan implementasi algoritma ke dalam bahasa pemrograman Go.

## Referensi
1. The Go Authors. (n.d.). A Tour of Go — Basics.https://go.dev/tour/basics
2. The Go Authors. (n.d.). A Tour of Go — Flow Control Statements: for, if, else. https://go.dev/tour/flowcontrol
3. The Go Authors. (n.d.). Package fmt Documentation. https://pkg.go.dev/fmt