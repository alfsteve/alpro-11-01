# <h1 align="center">Laporan Praktikum Modul 04 - Runtutan/Sekuensi</h1>
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

### 1. grade.go

```go
package main // Menentukan bahwa program ini berada di package main.
// Package main digunakan untuk membuat program Go yang bisa dijalankan.

import ( // Membuka bagian untuk mengimpor library/package yang dibutuhkan.
	"bufio" // Digunakan untuk membaca input dari keyboard dengan Scanner.
	"fmt"   // Digunakan untuk input dan output, seperti Print, Scan, dan Printf.
	"os"    // Digunakan untuk mengakses input/output dari sistem operasi.
)

func main() { // Fungsi utama program. Program mulai dijalankan dari sini.

	var nama string // Membuat variabel "nama" bertipe string untuk menyimpan nama.
	var nilai float64 // Membuat variabel "nilai" bertipe float64 untuk menyimpan nilai angka.
	var grade string // Membuat variabel "grade" bertipe string untuk menyimpan huruf nilai (A-F).

	scanner := bufio.NewScanner(os.Stdin) 
	// Membuat Scanner untuk membaca input dari keyboard.
	// os.Stdin berarti input berasal dari keyboard.

	fmt.Print("Masukkan nama: ")
	// Menampilkan teks "Masukkan nama: " ke layar.
	// Print tidak membuat baris baru.

	scanner.Scan()
	// Membaca input yang diketik oleh pengguna.
	// Hasil input sementara disimpan oleh scanner.

	nama = scanner.Text()
	// Mengambil teks yang sudah dibaca oleh scanner.
	// Kemudian memasukkannya ke dalam variabel nama.

	fmt.Print("Masukkan nilai: ")
	// Menampilkan teks "Masukkan nilai: " ke layar.

	fmt.Scan(&nilai)
	// Membaca angka yang dimasukkan pengguna.
	// &nilai berarti alamat memori variabel nilai diberikan kepada Scan
	// agar nilai yang dimasukkan dapat disimpan ke variabel nilai.

	if nilai >= 90 && nilai <= 100 {
		// Mengecek apakah nilai berada di antara 90 sampai 100.
		// && berarti "DAN", jadi kedua kondisi harus benar.

		grade = "A"
		// Jika kondisi benar, grade diisi dengan "A".

	} else if nilai >= 80 && nilai < 90 {
		// Jika kondisi sebelumnya salah, cek apakah nilai 80 sampai kurang dari 90.

		grade = "B"
		// Jika benar, grade diisi dengan "B".

	} else if nilai >= 70 && nilai < 80 {
		// Jika kondisi sebelumnya salah, cek apakah nilai 70 sampai kurang dari 80.

		grade = "C"
		// Jika benar, grade diisi dengan "C".

	} else if nilai >= 60 && nilai < 70 {
		// Jika kondisi sebelumnya salah, cek apakah nilai 60 sampai kurang dari 70.

		grade = "D"
		// Jika benar, grade diisi dengan "D".

	} else {
		// Jika semua kondisi di atas salah, jalankan bagian ini.

		grade = "F"
		// Grade diisi dengan "F".
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
	// Menampilkan nama dan grade ke layar.
	// %s digunakan untuk menampilkan data bertipe string.
	// \n digunakan untuk pindah ke baris baru.
	// Nilai pertama %s diambil dari variabel nama.
	// Nilai kedua %s diambil dari variabel grade.
}
```
#### Deskripsi
Program ini merupakan program yang berfungsi untuk menentukan grade nilai mahasiswa berdasarkan nilai angka yang dimasukkan oleh user. Program ini dmulai dengan package main sebagai package utama, yang kemudian mengimpor library bufio, fmt, dan os untuk membantu proses input dan output-nya. Di dalam func main(), terdapat tiga variabel, yaitu 'nama' bertipe string untuk menyimpan nama mahasiswa, 'nilai' bertipe float64 untuk menyimpan angka, dan 'grade' bertipe string untuk menyimpan hasil nilai huruf. Program menggunakan bufio.NewScanner(os.Stdin) untuk membaca input nama dari keyboard, sedangkan fmt.Scan(&nilai) digunakan untuk membaca nilai angka yang dimasukkan oleh pengguna. Setelah nama dan nilai diterima, program menggunakan percabangan if-else untuk menentukan grade berdasarkan rentang nilai. Nilai 90 sampai 100 mendapatkan grade A, nilai 80 hingga kurang dari 90 mendapatkan grade B, nilai 70 mendapatkan grade D. Sementara itu, nilai yang tidak memenuhi kondisi tersebut akan mendapatkan grade F. Operator perbandingan seperti >=, <=, dan < digunakan untuk memeriksa rentang nilai, sedangkan operator && itu digunakan untuk memastikan kalau kedua kondisi terpenuhi. Terakhir, program menggunakan fmt.Printf() untuk menampilkan nama mahasiswa beserta grade yang didapat. Dengan demikian, program ini menerapkan konsep variabel, input, dan output, serta percabangan if-else untuk menentukan nilai huruf mahasiswa secara otomatis.

### 2. penilaian.go

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
#### Deskripsi
Program ini digunakan untuk membuat sistem penilaian siswa berbasis menu. Program dimulai dengan library bufio, fmt, dan os, serta mendefinisikan variabel 'pilihan,' 'nama,' 'nilai,' dan 'grade' sesuai tipe datanya. Selanjutnya, Program menampilkan menu yang menyediakan pilihan 1 untuk sistem penilaian dan 0 untuk keluar. Jika pengguna memilih 1, program akan meminta nama dan nilai siswa, kemudian menentukan grade dengan menggunakan percabangan if-else, yaitu A untuk nilai 90-100, B untuk 80 hingga kurang dari 90, C untuk 70 hingga kurang dari 80, D untuk 60 hingga kurang dari 70, dan F untuk nilai lainnya. Hasil penilaian kemudian ditampilkan menggunakan fmt.Printf(). Jika pengguna memilih 0, program menampilkan pesan bahwa program telah keluar, sedangkan pilihan selain 1 dan 0 akan menampilkan pesan bahwa pilihan tidak valid. Dengan demikian, program ini menerapkan konsep variabel, input, dan output, serta percabangan if-else untuk membuat sistem penilaian sederhana berbasis menu.

### 3. klasifikasi.go

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}

```
#### Deskripsi
Program ini digunakan untuk mengklasifikasikan seseorang berdasarkan usia dan gaji. Program mengimpor library fmt untuk membaca input dan menampilkan output, kemudian mendefinisikan variabel 'usia' dan 'gaji' bertipe int, serta 'keterangan' bertipe string. Pengguna memasukkan usia dan gaji, lalu program menggunakan percabangan switch tanpa ekspresi untuk memeriksa kondisi secara berurutan. Jika usia di bawah 18 tahun, keterangannya adalah "Masih sekolah". Untuk usia 18-25 tahun, gaji minimal 50 menghasilkan keterangan "Muda sukses", sedangkan gaji di bawah 50 menghasilkan keterangan "Masih belajar hidup". Untuk usia 26-40 tahun, gaji minimal 100 dikategorikan sebagai "Pekerja mapan", sedangkan gaji di bawah 100 sebagai "Perlu perbaikan karier". Sementara itu, usia di atas 40 tahun dengan gaji minimal 150 dikategorikan sebagai "Profesional berpengalaman", sedangkan gaji di bawah 150 sebagai "Perlu evaluasi finansial". Terakhir, program menggunakan fmt.Println() untuk menampilkan hasil klasifikasi. Dengan demikian, program ini menerapkan konsep variabel, input-output, operator perbandingan, operator logika, dan percabangan switch untuk menentukan keterangan berdasarkan kondisi usia dan gaji.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. modif.go

```go
package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	switch nilai {
		case nilai >= 90 && nilai <= 100:
			grade = "A"
		case nilai >= 80 && nilai < 90:
			grade = "B"
		case nilai >= 70 && nilai < 80:
			grade = "C"
		case nilai >= 60 && nilai < 70:
			grade = "D"
		default:
			grade = "E"
	}
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Grade: %s\n", grade)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/modif_grade/output.png)


#### Deskripsi
Program ini digunakan untuk menentukan grade berdasarkan nama dan nilai yang dimasukkan pengguna. Program menggunakan 'bufio.Scanner' untuk membaca nama, 'fmt.Scan' untuk membaca nilai, dan 'switch' untuk menentukan grade A sampai E sesuai rentang nilai. Jika nilai di luar rentang 0–100, program menampilkan keterangan bahwa nilai tidak valid. Terakhir, program menampilkan nama dan grade pengguna.

### 2. pajak.go

```go

package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	fmt.Print("Masukkan penghasilan (juta): ")
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 5.0 / 100 * penghasilan
	} else if penghasilan <= 100 {
		pajak = (5.0/100 * 50) +
			(10.0/100 * (penghasilan - 50))
	} else if penghasilan <= 200 {
		pajak = (5.0/100 * 50) +
			(10.0/100 * 50) +
			(15.0/100 * (penghasilan - 100))
	} else {
		pajak = (5.0/100 * 50) +
			(10.0/100 * 50) +
			(15.0/100 * 100) +
			(20.0/100 * (penghasilan - 200))
	}

	fmt.Printf("Total pajak: %g juta\n", pajak)
}

```

##### Output
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/pajak/output.png)

#### Deskripsi
Program ini digunakan untuk menghitung pajak berdasarkan total penghasilan yang dimasukkan pengguna. Program menggunakan variabel penghasilan untuk menyimpan jumlah penghasilan dan pajak untuk menyimpan hasil perhitungan. Dengan menggunakan if-else, program menentukan tarif pajak sesuai rentang penghasilan, mulai dari 5% hingga 20%. Pajak dihitung secara bertingkat berdasarkan setiap lapisan penghasilan. Terakhir, program menampilkan total pajak yang harus dibayar.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan program yang telah dibuat, dapat disimpulkan bahwa penyelesaian masalah dalam praktikum ini dilakukan dengan menentukan input, proses, dan output terlebih dahulu. Penggunaan variabel, tipe data, serta struktur percabangan membantu program mengolah data dan menghasilkan output sesuai dengan kondisi yang diberikan. Program-program tersebut menerapkan konsep dasar analisis masalah dan implementasi algoritma ke dalam bahasa pemrograman Go.

## Referensi
1. The Go Authors. (n.d.). A Tour of Go. https://go.dev/tour/
2. The Go Authors. (n.d.). Go Documentation. https://go.dev/doc/
3. The Go Authors. (n.d.). fmt package. https://pkg.go.dev/fmt
4. The Go Authors. (n.d.). bufio package. https://pkg.go.dev/bufio