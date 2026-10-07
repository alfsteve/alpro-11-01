# <h1 align="center">Tugas Pendahuluan Modul [04] - [RUNTUTAN DAN SEKUENSI]</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

```go
package main
import "fmt"

func main() {
	var intNumber, a int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a)
	intNumber = 5
	if a < intNumber {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}


}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal1/output.png)


#### Deskripsi
[Program ini digunakan untuk menentukan nilai dari True dan False suatu angka. Di sini terdapat dua variabel dengan nama intNumber dan a sebagai integer, lalu di sini variabel a menggunakan fmt.Scan agar nilainya bisa diinput sendiri. Setelah itu variabel intNumber itu bernilai 5, dan jika variabel a lebih dari angka 5 maka false, jika variabel a kurang dari angka 5 maka dia akan menjadi true.]

### 2. Tracing

```go
 package main 
  
  import "fmt" 
   
  func main() { 
      x := 10 
      y := 5 
      z := 15 
      result := 0 
      if x > 5 { 
          if y < 10 { 
              result = x + y 
          } else { 
              result = x - y 
          } 
      } 
   
      if z > 10 && x == 10 { 
          result += z 
     	 } else { 
          result = z - x 
      } 
   
      if x == 10 || y > 10 { 
          result += 5 
     	 } else if y == 5 && z > 10 { 
          result -= 5 
     	 } else { 
          result *= 2 
      } 
   
      if !(x < 15 && y < 10) { 
          result += 10 
     	 } else { 
          result -= 10 
      }    
		fmt.Println("Nilai akhir result:", result)  
	} 

```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal2/output.png)


#### Deskripsi
[Program ini digunakan untuk menghitung nilai dari x, y, z. Dengan variabel x yang bernilai 10, variabel y yang bernilai 5, dan variabel z yang bernilai 15. Untuk hasil dari setiap pertanyaan nomor 2 ini sebagai berikut:
1. Nilai akhir dari variabel result ini adalah 25.
2. Outputnya berupa: "Nilai akhir result: 25"
3. Langkah-langkah dari alur eksekusi programnya sebagai berikut
	-Kondisi x > 5 itu benar, yang terjadi setelahnya adalah pernyataan kondisi selanjutnya yaitu jika y < 10, lalu jika kondisinya benar maka result = x + y.
	-Kondisi z > 10 && x == 10 itu benar, karena kondisi tersebut benar maka pernyataan kondisi selanjutnya yaitu result += z, yang artinya result dari x + y tadi ditambah dengan variabel z, itulah hal yang mempengaruhi nilai result.
	-Salah satu kondisi kondisi yang benar yaitu x == 10, yang terjadi selanjutnya yaitu result += 5, yang artinya result hasil dijumlahkan dengan z tadi dijumlahkan lagi dengan 5.]


### 3. Jumlah Hari dalam sebulan berdasarkan Tahun dan Bulan

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal3/output.png)


#### Deskripsi
[Program ini digunakan digunakan untuk menghitung jumlah hari dalam sebulan berdasarkan input tahun dan nama bulan dari pengguna, alurnya itu melalui pernyataan kondisi if bulan Jan, Mar, Mei, Jul, Aug, Okt, dan Des itu 31 hari, lalu pernyataan kondisi else if 30 hari yang meliputi bulan Apr, Jun, Sep dan Nov, dan pernyataan kondisi else if Feb yang 28 hari, lalu jika program mendeteksi tahun kabisat maka program akan mencetak tambahan angka 29.]

### 4. Switch Case

```go
package main
import "fmt"

func main() {
	var warna string
	fmt.Print("Masukkan warna: ")
	fmt.Scan(&warna)

	switch warna {
		case "Merah":
			fmt.Println("Aku marahhh >:(")
		case "Kuning":
			fmt.Println("Aku sedang khawatir :(")
		case "Hijau":
			fmt.Println("Aku merasa tenang :)")
		default:
			fmt.Println("Warna tidak dikenali mas")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal4/output.png)


#### Deskripsi
[Program ini digunakan untuk memberikan kondisi perasaan saat pengguna memasukkan suatu warna. Pada program terdapat variabel warna yang dinyatakan sebagai string, lalu variabel warna tersebut dijadikan input menggunakan fmt.Scan(&warna), setelah itu menggunakan pernyataan kondisi switch case, dan di dalamnya terdapat 3 case yang berisi Merah, Kuning, Hijau. Setelah itu di case merah itu mencetak perasaan marah, kuning itu khawatir, dan hijau itu tenang.]

## Kesimpulan
[Kesimpulan yang didapat dari pengerjaan laporan tugas pendahuluan modul 4 ini adalah kita bisa mengetahui fungsi-fungsi dari penggunaan if else dan switch case dan banyak cara yang bisa digunakan untuk memecahkan berbagai tugas yang akan diberikan kedepannya.]