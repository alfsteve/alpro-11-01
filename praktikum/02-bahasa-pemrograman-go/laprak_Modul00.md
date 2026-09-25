# <h1 align="center">Laporan Praktikum Modul 02 - Variabel, Tipe Data, dan Operasi</h1>
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
Digunakan untuk menyimpan bilangan bulat. Variabelnya berupa "int, int8, int32, int64, uint, uint8, uint32, dan uint64."

##### b. Real 
Digunakan untuk menyimpan bilangan pecahan atau bilangan real. Variabelnya berupa "float32 dan float64."

##### c. Boolean
Digunakan untuk menyimpan nilai true atau false. Variabelnya berupa "bool."

##### d. String
Digunakan untuk menyimpan kumpulan karakter atau teks.


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. suhu.go

```go
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
```
#### Deskripsi
Program ini digunakan untuk mengonversi suhu Celcius ke Reamur, Fahrenheit, dan Kelvin. Program ini menerima input dari suhu Celcius yang kemudian masing-masing hasil dihitung menggunakan rumus konversi. Hasil konversi tersebut kemudian ditampilkan dalam satu baris dengan satuan Celcius, Reamur, Fahrenheit, dan Kelvin.

### 2. lingkaran.go

```go
package main

import "fmt"

func main() {

	//Membuat variabel
	var pi float64 = 3.14
	var r float64

	//Membaca input
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	//Membaca Output
	fmt.Print("Luas lingkaran: ", pi, " * ", r, " * ", r, " = ")
	fmt.Println(pi * r * r)
}
```
#### Deskripsi
Program ini digunakan untuk menghitung luas lingkaran berdasarkan jari-jari yang dimasukkan. Program menggunakan variabel pi dan r, kemudian menghitung luas dengan rumus pi * r * r. Hasil perhitungan luas lingkaran setelah input dimasukkan dan diproses.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. kalkulator.go

```go
package main

	import "fmt"
	
	func main() {
		var a, b int

		//Membaca input
		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan nilai b: ")
		fmt.Scan(&b)

		//Menampilkan output
		fmt.Print("Hasil Penjumlahan: ")
		fmt.Println(a + b)
		fmt.Print("Hasil Pengurangan: ")
		fmt.Println(a - b)
		fmt.Print("Hasil Perkalian: ")
		fmt.Println(a * b)
		fmt.Print("Hasil Pembagian: ")
		fmt.Println(a / b)
		fmt.Print("Hasil Sisa Hasil Bagi: ")
		fmt.Println(a % b)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/kalkulator/output.png)


#### Deskripsi
Program ini digunakan untuk melakukan operasi aritmatika pada dua bilangan, yaitu penjumlahan, pengurangan, perkalian, pembagian, dan sisa hasil bagi. Program menerima inputan dari nilai a dan b, kemudian menghitung lalu menampilkan hasi dari setiap operasinya.

### 2. cacahuang.go

```go
package main

	import "fmt"

	func main() {
		var a, b, c int

		//Membaca input
		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan nilai b: ")
		fmt.Scan(&b)
		fmt.Print("Masukkan nilai c: ")
		fmt.Scan(&c)

		//Menampilkan output
		fmt.Print("Hasilnya : ", a, " * 10000 + ", b, " * 5000 + ", c, " * 1000 = ")
		fmt.Println(a * 10000 + b * 5000 + c * 1000)
		
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
1. The Go Authors. (2026). The Go Programming Language Specification. Diakses pada 25 September 2026 melalui https://go.dev/ref/spec
2. The Go Authors. (2026). Documentation - The Go Programming Language. Diakses pada 25 September 2026 melalui https://go.dev/doc/
