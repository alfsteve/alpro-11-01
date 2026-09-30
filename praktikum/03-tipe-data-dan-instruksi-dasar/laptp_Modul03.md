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
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/sisa/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. Boolean

```go
package main

	import "fmt"
	
	func main(){

		var a bool
		var b bool

		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		

		fmt.Println("nilai dari a adalah", a)


}

```

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
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/konversi/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]