# <h1 align="center">Tugas Pendahuluan Modul [03] - [VARIABEL DAN OPERATOR]</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

### 1. Sisa Kue

```go
package main
import "fmt"

func main() {
	var y, x int

	fmt.Print("Masukkan jumlah kue: ")
	fmt.Scan(&y)

	fmt.Print("Masukkan jumlah anggota keluarga: ")
	fmt.Scan(&x)

	fmt.Print("Jumlah kue yang tersisa: ")
	fmt.Println(y % x)
	
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/kue/Output.png)


#### Deskripsi
[Program ini digunakan untuk menghitung sejumlah kue yang tersisa setelah sejumlah kue itu dibagikan kepada beberapa anggota keluarga secara merata. Program ini menerima dua input, yaitu jumlah kue dan jumlah anggota keluarga. Operator yang digunakan itu %, operator ini digunakan untuk mendapatkan sisa hasil pembagian jumlah kue dengan jumlah anggota keluarga.]

### 2. Boolean

```go
package main

	import "fmt"
	
	func main(){

		var a bool

		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		

		fmt.Println("nilai dari a adalah", a)


}

```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/boolean/Output.png)


#### Deskripsi
[Program ini digunakan untuk menerima dan menampilkan sebuah nilai boolean, yaitu nilai yang hanya memiliki dua kemungkinan, yaitu true atau false. Program ini menggunakan variabel a dengan tipe data bool]


### 3. Mil ke Kilo

```go
package main
import "fmt"
func main() {

	var mil float64

	fmt.Print("Masukkan jumlah mil: ")
	fmt.Scan(&mil)

	fmt.Printf("Hasil konversinya sebagai berikut: %.1f km", mil * 1.6)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/mil_ke_kilo/Output.png)


#### Deskripsi
[Program ini digunakan untuk mengkonversi jarak dari satuan mil (mile) ke kilometer (km). Program menerima jumlah mil dari pengguna, kemudian mengalikannya dengan nilai konversi 1 mil = 1,6 kilometer.]

## Kesimpulan
[Kesimpulan yang di dapat dari pembuatan ketiga program ini yaitu, kita dilatih untuk menggunakan logika untuk mengikuti soal-soal dari tugas tersebut, dan membiasakan diri untuk mengetik dan mengingat kode-kode yang digunakan.]