# <h1 align="center">Laporan Praktikum Modul 03 - Variabel, Tipe Data, dan Operasi</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

## Dasar Teori

### A. Bahasa Pemrograman Go (Golang)
Go atau Golang merupakan bahasa pemrograman yang digunakan dalam praktikum Algoritma Pemrograman. Pada modul 02 ini, pembahasannya berfokus pada penggunaan bahasa Go untuk membuat suatu program yang melibatkan variabel, tipe, data, dan operasi.

### B. Struktur Pemrograman Go

#### 1. Pengertian package main dan func main()
package main merupakan penanda kalau sebuah file berisi program utama dalam bahasa Go. Sementara itu, func main() berisi kode atau instruksi yang akan dijalankan ketika program dieksekusi. Kedua komponen ini merupakan bagian dasar dari struktur dalam program Go.

#### 2. import, fmt, Komentar, dan Variabel
import digunakan untuk memasukkan package yang diperlukan dalam program. Package fmt digunakan untuk proses input dan output, seperti fmt.Scanln() untuk membaca masukan (Input) dan fmt.Println() atau fmt.Print() untuk menampilkan keluaran (Output). Komentar dapat ditulis menggunakan // untuk satu baris atau /* ... */ untuk beberapa baris. Selain itu, variabel dapat dideklarasikan menggunakan var untuk menyimpan data dengan tipe tertentu.

#### 3. Struktur Dasar Pemgrograman Go
Variabel merupakan nama dari suatu lokasi di memori yang digunakan untuk menyimpan data dengan tiper tertentu. Dalam Go, variabel dapat dideklarasikan menggunakan var. Beberapa tipe data dasar yang terdapat dalam modul meliputi:

##### a. Integer
Adalah tipe data yang digunakan untuk menyimpan bilangan bulat, yaitu bilangan yang tidak memiliki angka pecahan atau koma. Dalam Go, tipe integer itu terdiri dari int, int8, int32, int64, uint, uint8, uint32, dan uint64. Tipe ini dapat digunakan untuk menyimpan nilai seperti jumlah, skor, atau angka lainnya yang merupakan bilangan bulat.

##### b. Real 
Adalah tipe data yang digunakan untuk menyimpan bilangan pecahan atau bilangan real, yaitu bilangan yang dapat memiliki angka di belakang koma. Dalam Go, tipe data real terdiri dari float32 dan float64. Tipe ini digunakan ketika program membutuhkan nilai yang lebih tepat dalam bentuk pecahan atau hasil perhitungan desimal.

##### c. Boolean
Adalah tipe data yang digunakan untuk menyimpan nilai logika, yaitu true atau false. Dalam Go, tipe data Boolean menggunakan bool. Nilai Boolean ini biasanya digunakan untuk menunjukkan suatu kondisi yang bernilai benar atau salah.

##### d. String
Adalah tipe data yang digunakan untuk menyimpan kumpulan karakter atau teks. Dalam Go, tipe data string menggunakan string. Tipe ini biasanya digunakan untuk menyimpan data berupa nama, kata, maupun kalimat yang diperlukan dalam program.


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. kasir.go

```go
package main
import "fmt"
func main(){

	var uang int
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&uang)

	var sepuluhRibu int = uang / 10000
	var sisa int = uang % 1

	var limaRibu int = uang / 5000
	sisa = sisa % 5000

	var seRibu int = uang / 1000
	
	fmt.Println(sepuluhRibu, limaRibu, seRibu)
}
```
#### Deskripsi
Program ini digunakan untuk mengonversi suhu Celcius ke Reamur, Fahrenheit, dan Kelvin. Program ini menerima input dari suhu Celcius yang kemudian masing-masing hasil dihitung menggunakan rumus konversi. Hasil konversi tersebut kemudian ditampilkan dalam satu baris dengan satuan Celcius, Reamur, Fahrenheit, dan Kelvin.

### 2. konversi.go

```go
package main
	import "fmt"
	func main() {

		var c float32

		fmt.Print("Masukkan suhu celcius :")
		fmt.Scan(&c)

		fmt.Println(c + 273)

	}
```
#### Deskripsi
Program ini digunakan untuk menghitung luas lingkaran berdasarkan jari-jari yang dimasukkan. Program menggunakan variabel pi dan r, kemudian menghitung luas dengan rumus pi * r * r. Hasil perhitungan luas lingkaran setelah input dimasukkan dan diproses.

### 3. tukar.go

```go
package main
import "fmt"
func main(){

	var x, y, z int
	
	fmt.Scan(&x, &y, &z)
	temp := x
	z = y
	y = temp

	fmt.Println(x, y, z)
}

```
#### Deskripsi
Program ini digunakan untuk menghitung total dan rata-rata skor dari dua mata pelajaran, yaitu Emteka dan Bahasa Inggris. Program menerima input nama dan kedua skor, kemudian menjumlahkan skor untuk mendapatkan total dan membaginya dengan 2 untuk mendapatkan rata-rata. Hasil berupa nama, total skor, dan rata-rata skor kemudian ditampilkan sebagai output.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. hari.go

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/kalkulator/output.png)


#### Deskripsi
Program ini digunakan untuk melakukan operasi aritmatika pada dua bilangan, yaitu penjumlahan, pengurangan, perkalian, pembagian, dan sisa hasil bagi. Program menerima inputan dari nilai a dan b, kemudian menghitung lalu menampilkan hasi dari setiap operasinya.

### 2. reamur.go

```go
package main
import "fmt"
func main(){

	var c, reamur float32
	fmt.Print("Masukkan angka Celcius: ")
	fmt.Scan(&c)

	reamur = c * 4.0 / 5.0

	fmt.Println("Reamur =", reamur)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)

#### Deskripsi
Program ini digunakan untuk menghitung total nilai uang berdasarkan jumlah pecahan Rp10.000, Rp5.000, dan Rp1.000 yang dimasukkan. Program menerima input jumlah masing-masing pecahan, kemudian menghitung total nilainnya menggunakan operasi perkalian dan penjumlahan. Hasil total nilai uang ditampilkan setelah proses menghitung.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Praktikum Modul 02 ini memberikan pemahaman pada saya tentang variabel, tipe data, input, output, operasi dasar dalam bahasa pemrograman Go. Melalui tugas guided dan unguided ini, konsepnya akan diterapkan dalam program perhitungan nilai, pertukaran variabel, konversi suhu, luas lingkaran, operasi aritmatika, dan perhitungan nilai uang. Hasil praktikum menunjukkan bahwa konsep dasar tersebut dapat digunakan untuk membuat program sederhana yang menerima input, melakukan proses, dan menghasilkan output sesuai kebutuhan kita.

## Referensi
1. The Go Authors. (2025). The Go Programming Language Specification. California: Google Inc. Diakses pada 27 September 2026 melalui https://go.dev/ref/spec
2. The Go Authors. (2025). Documentation - The Go Programming Language. California: Google Inc. Diakses pada 27 September 2026 melalui https://go.dev/doc/