package main // Menentukan bahwa program menggunakan package main.
// Package main digunakan untuk membuat program yang bisa dijalankan.

import "fmt" // Mengimpor package fmt untuk input dan output.

func main() { // Fungsi utama. Program mulai dijalankan dari sini.

	var tahun int // Membuat variabel tahun bertipe int untuk menyimpan angka tahun.
	var bulan string // Membuat variabel bulan bertipe string untuk menyimpan nama bulan.

	fmt.Print("Masukkan tahun: ")
	// Menampilkan teks "Masukkan tahun: " ke layar.

	fmt.Scan(&tahun)
	// Membaca angka tahun yang dimasukkan pengguna.
	// &tahun digunakan agar input disimpan ke variabel tahun.

	fmt.Print("Masukkan bulan: ")
	// Menampilkan teks "Masukkan bulan: " ke layar.

	fmt.Scan(&bulan)
	// Membaca nama bulan yang dimasukkan pengguna.
	// Input disimpan ke variabel bulan.

	if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" || bulan == "Jul" || bulan == "Aug" || bulan == "Okt" || bulan == "Des" {
		// Mengecek apakah bulan adalah salah satu dari:
		// Jan, Mar, Mei, Jul, Aug, Okt, atau Des.
		//
		// == berarti "sama dengan".
		// || berarti "ATAU".
		// Jadi cukup salah satu kondisi benar untuk masuk ke bagian ini.

		fmt.Println(31)
		// Menampilkan angka 31.
		// Artinya bulan tersebut memiliki 31 hari.

	} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov" {
		// Jika kondisi pertama salah, program mengecek apakah bulan:
		// Apr, Jun, Sep, atau Nov.
		//
		// Semua bulan tersebut memiliki 30 hari.

		fmt.Println(30)
		// Menampilkan angka 30.
		// Artinya bulan tersebut memiliki 30 hari.

	} else if bulan == "Feb" {
		// Jika kondisi sebelumnya salah, program mengecek apakah bulan adalah Februari.

		fmt.Println(28)
		// Menampilkan 28 sebagai jumlah hari Februari pada kondisi normal.

		if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
			// Mengecek apakah tahun tersebut merupakan tahun kabisat.
			//
			// % adalah operator modulus, yaitu mencari SISA hasil pembagian.
			//
			// tahun%4 == 0
			// Artinya tahun habis dibagi 4.
			//
			// tahun%100 != 0
			// Artinya tahun TIDAK habis dibagi 100.
			//
			// tahun%400 == 0
			// Artinya tahun habis dibagi 400.
			//
			// && berarti DAN.
			// || berarti ATAU.
			//
			// Aturan tahun kabisat:
			// Tahun harus habis dibagi 4
			// DAN
			// (tidak habis dibagi 100 ATAU habis dibagi 400).

			fmt.Println(29)
			// Jika tahun merupakan tahun kabisat,
			// Februari memiliki 29 hari.
		}
	}
}