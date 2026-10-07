# <h1 align="center">Laporan Praktikum Modul 03 - Variabel dan Operator</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

## Dasar Teori

### A. Bahasa Pemrograman Go (Golang)
Dalam pemrograman ini, penyelesaian masalahnya dilakukan dengan mengidentifikasi data masukan (input), proses yang dilakukan, dan informasi keluaran (output). Selain itu, tipe data harus ditentukan sesuai dengan data yang digunakan, seperti int untuk bilangan bulat dan float32 untuk bilangan rill.

Bahasa Go menggunakan package main, import, dan func main() sebagai struktur dasar program. Input dapat dibaca menggunakan fmt.Scan(), sedangkan output dapat ditampilkan menggunakan fmt.Println(). Operator aritmatika seperti +, *, /, dan % digunakan dalam proses perhitungan.

### B. Struktur Pemrograman Go

#### 1. Input
Menerima data dari pengguna menggunakan fmt.Scan().

#### 2. Proses
Melakukan perhitungan menggunakan operator aritmatika dan rumus yang diperlukan.

#### 3. Output
Menampilkan hasil menggunakan fmt.Println().

Program juga menggunakan variabel dengan tipe data yang sesuai. Contohnya, int digunakan untuk jumlah hari dan uang, sedangkan float32 digunakan untuk nilai suhu yang dapat memiliki angka desimal. Penentuan tipe data merupakan bagian dari analisis masalah sebelum algoritma diterjemahkan ke bahasa pemrograman.


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
Program ini digunakan untuk menghitung suhu melalui konversi dari celcius. Disini variabel dari c sebagai float32, lalu diinput suhu celcius tersebut melalui fmt.Scan dan keluarannya berupa inputan angka dari celcius kemudian ditambahkan angka 273 sehingga menghasilkan suhu yang dikonversi.

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
Program ini digunakan untuk menukar 3 variabel berupa x, y, dan Z. Variabel x, y, z sebagai integer (int). Kemudian ketiga variabel tersebut menggunakan fmt.Scan agar bisa diinput sendiri dan kemudian x disimpan ke temp untuk sementara, lalu z ditukar ke y dan y tadi ditukar ke temp yang awalnya x. Lalu keluaran dari nilai x, y, z yang ditukar tadi ditampilkan.

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
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/hari/Output.png)


#### Deskripsi
Program ini digunakan untuk mengubah jumlah hari menjadi tahun, bulan, minggu, dan hari. Program menggunakan pembagian (/) untuk mendapatkan hasil dan modulus (%) untuk mendapatkan sisa hari. Asumsi yang digunakan adalah 1 tahun = 360 hari, 1 bulan = 30, dan 1 minggu = 7 hari.

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
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/reamur/Output.png)

#### Deskripsi
Program ini digunakan untuk mengubah suhu dari Celsius ke Reamur. Nilai Celsius dimasukkan oleh pengguna, kemudian dikonversi menggunakan rumus Reamur = Celsius * 4 / 5, lalu hasilnya ditampilkan

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan keempat program, dapat disimpulkan bahwa praktikum modul 03 ini penyelesaian masalahnya dengan menentukan input, proses, dan output-nya terlebih dahulu. Pemilihan tipe data yang tepat serta penggunaan operator aritmatika membantu program menghasilkan output sesuai dengan permasalahan yang diberikan. Keempat program menerapkan konsep dasar analisis masalah dan implementasi algoritma ke dalam bahasa Go.

## Referensi
1. Go. (2026). The Go Programming Language. https://go.dev/
2. Go Documentation. (2026). The Go Programming Language Documentation. https://go.dev/doc/
3. Go Package Documentation. (2026). fmt Package. https://pkg.go.dev/fmt