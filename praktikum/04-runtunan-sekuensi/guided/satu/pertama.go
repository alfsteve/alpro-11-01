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