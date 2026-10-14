# <h1 align="center">Tugas Pendahuluan Modul [05] - [PERULANGAN FOR]</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

### 1. Fungsi Utama Perulangan for dalam Go

Perulangan for dalam bahasa pemrograman Go ini berfungsi untuk menjalankan suatu blok kode secara berulang selama kondisi tertentu terpenuhi. Perulangan ini membantu mengurangi penulisan kode yang sama secara berulang.

Perulangan for memiliki tiga komponen utama, yaitu:

#### a. Inisialisasi

Menentukan nilai awal variabel penghitung sebelum perulangannya dimulai. Contohnya, i := 1.

#### b. Kondisi

Menentukan apakah perulangan akan dilanjutkan atau dihentikan. Contohnya, i <= 5.

#### c. Post Statement

Memperbarui nilai variabel setelah setiap putaran perulangan. Contohnya, i++, yang artinya nilai i bertambah satu.

### 2. Perbedaan for dengan tiga Komponen dan for dengan Kondisi saja

#### a. for dengan tiga Komponen

Bentuk ini menggunakan inisialisasi, kondisi, dan post statement. Perulangan ini cocok digunakan ketika variabel penghitung perlu diinisiasi dan diperbarui pada setiap putaran.

```go
 package main 
 import "fmt" 
 func main() { 

     for i:= 1; i <= 5; i++ {
		fmt.Println(i)
	 }

}

```

#### Deskripsi
Pada program ini, variabel i diinisialisasi sebagai 1, perulangan berlangsung selama i <= 5, dan nilai i bertambah satu setelah setiap putaran.

#### b. for dengan Kondisi saja

Bentuk ini hanya menggunakan kondisi pada bagian for, sehingga cara kerjanya itu mirip dengan perulangan while pada bahasa lain. Inisialisasi dan perubahan nilai variabel dapat ditulis di dalam blok perulangan.

```go
package main
import "fmt"
func main(){
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
```

#### Deskripsi
Pada program ini, perulangan berlangsung selama nilai i tidak lebih dari 5. Perintah i++ menaikkan nilai i setiap putaran sehingga perulangan akhirnya berakhir.


### 3. ganjil.go

```go
package main
import "fmt"
func main(){

	var n int
	var a int = 0
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		a += i

	}
	fmt.Println()
	fmt.Println("hasil :", a)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal3/output.png)


#### Deskripsi
Program ganjil.go ini digunakan untuk menampilkan seluruh bilangan ganjil dari 1 hingga batas bilangan bulat positif n yang dimasukkan oleh pengguna, serta menghitung jumlah seluruh bilangan ganjil tersebut. Program ganjil.go ini menggunakan variabel n sebagai batas bilangan dan variabel jumlah untuk menyimpan hasil penjumlahan. Perulangan for dimulai dari angka 1 lalu bertambah 2 pada setiap iterasi sehingga hanya bilangan ganjil yang ditampilkan. Pada setiap iterasi, nilai bilangan ganjil ditambahkan ke variabel jumlah. Setelah perulangannya selesai, program menampilkan total penjumlahan seluruh bilangan ganjil yang diperoleh.

### 4. perkalian.go

```go
package main
import "fmt"
func main(){

	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)


	for i := 1; i <= b; i++ {
	
		hasil := a * i
		fmt.Println(a, "x", i, "=", hasil)

	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/soal4/output.png)


#### Deskripsi
Program perkalian.go digunakan untuk menampilkan tabel perkalian dari bilangan a mulai dari a × 1 hingga a × b, berdasarkan dua bilangan bulat positif yang dimasukkan oleh pengguna. Program menggunakan variabel a sebagai bilangan yang dikalikan dan variabel b sebagai batas perkalian. Perulangan for dimulai dari angka 1 dan bertambah 1 pada setiap iterasi hingga mencapai nilai b. Pada setiap iterasi, program menghitung hasil perkalian antara a dan variabel penghitung i, kemudian menampilkan hasilnya dalam format operasi perkalian. Dengan demikian, program menghasilkan sejumlah baris sesuai dengan nilai b yang dimasukkan.

## Kesimpulan
Berdasarkan kedua program dan soal di atas, dapat disimpulkan bahwa perulangan for dalam bahasa Go dapat digunakan untuk menyelesaikan permasalahan yang membutuhkan proses berulang secara terstruktur. Pada program ganjil.go, perulangan digunakan untuk menghasilkan bilangan ganjil dengan penambahan nilai penghitung sebesar 2 sekaligus menghitung jumlahnya. Sementara itu, pada program perkalian.go, perulangan digunakan untuk menghasilkan tabel perkalian dengan penambahan nilai penghitung sebesar 1. Kedua program menunjukkan bahwa penggunaan variabel, operator aritmatika, dan perulangan for yang tepat dapat membantu menyelesaikan permasalahan secara sistematis dan menghasilkan keluaran sesuai dengan kebutuhan.