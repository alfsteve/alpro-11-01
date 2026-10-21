# <h1 align="center">Tugas Pendahuluan Modul [06] - [PERULANGAN WHILE]</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

### 1. lulus.go

```go
package main
import "fmt"
func main(){
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	for nilai != -1{
	if nilai >= 60{
		fmt.Println("Anda Lulus Ujian")
	} else {
		fmt.Println("Anda tidak Lulus")
	}
		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&nilai)
	}
	fmt.Println("Program selesai.")
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/lulus/output.png)

#### Deskripsi
Program ini digunakan untuk menentukan kelulusan siswa berdasarkan nilai ujian yang dimasukkan oleh pengguna. Program membaca nilai secara berulang menggunakan perulangan for dengan kondisi nilai != -1. Jika nilai lebih besar atau sama dengan 60, program menampilkan pesan "Anda lulus ujian." Namun, jika nilai kurang dari 60, program menampilkan pesan "Anda tidak lulus ujian." Proses akan berlangsung sampai pengguna memasukkan nilai -1, kemudian program menampilkan pesan "Program selesai." Program ini menerapkan konsep perulangan, percabangan if-else, input dan penggunaan nilai sentinel sebagai kondisi berhenti.

### 2. deret.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	
	i := 1
	jumlah := 0

	for i <= n {
		jumlah = jumlah + i
		i++
	}
	fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/deret/output.png)

#### Deskripsi
Program ini digunakan untuk menghitung jumlah deret bilangan dari 1 sampai bilangan bulat positif n. Program menggunakan variabel penghitung untuk menentukan angka yang sedang diproses dan variabel penjumlahan untuk menyimpan total sementara. Melalui perulangan for, setiap angka ditambahkan ke total hingga penghityng mencapai nilai n. Setelah kondisi perulangan tidak lagi terpenuhi, program menampilkan hasil penjumlahan. Program ini menerapkan konsep perulangan, variabel penghitung, dan akumulasi nilai.


### 3. duaterendah.go

```go
package main
import "fmt"
func main() {
    var n int
    fmt.Print("Jumlah mahasiswa: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Println("Jumlah mahasiswa minimal 2.")
    } else {
        var nama1, nama2 string
        var nilai1, nilai2 int

        fmt.Print("Nama dan nilai mahasiswa 1: ")
        fmt.Scan(&nama1, &nilai1)

        fmt.Print("Nama dan nilai mahasiswa 2: ")
        fmt.Scan(&nama2, &nilai2)

        if nilai2 < nilai1 {
            nama1, nama2 = nama2, nama1
            nilai1, nilai2 = nilai2, nilai1
        }

        var i int = 3
        for i <= n {
            var nama string
            var nilai int

            fmt.Print("Nama dan nilai mahasiswa ", i, ": ")
            fmt.Scan(&nama, &nilai)

            if nilai < nilai1 {
                nama2 = nama1
                nilai2 = nilai1
                nama1 = nama
                nilai1 = nilai
            } else if nilai < nilai2 {
                nama2 = nama
                nilai2 = nilai
            }

            i++
        }

        fmt.Println("Nilai terendah pertama:", nama1, nilai1)
        fmt.Println("Nilai terendah kedua:", nama2, nilai2)
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/duaterendah/output.png)

#### Deskripsi
Program ini digunakan untuk mencari nama siswa yang memperoleh nilai terendah dan nilai kedua terendah dari sejumlah siswa yang dimasukkan oleh pengguna. Program membaca jumlah siswa terlebih dahulu, kemudian menerima nama dan nilai setiap siswa. Dua nilai terendah disimpan dalam variabel, lalu dibandingkan dengan nilai berikutnya menggunakan percabangan if-else. Jika ditemukan nilai yang lebih rendah, program memperbarui nilai dan nama siswa pada posisi yang sesuai. Setelah semua siswa diproses, program menampilkan nama siswa dengan nilai terendah dan kedua terendah. Program ini menerapkan konsep perulangan, percabangan, perbandingan nilai, dan pembaruan variabel.

### 4. biaya.go

```go

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    biaya := 10000000.0
    tahun := 1

    for tahun < n {
        biaya = biaya * 1.06
        tahun++
    }

    fmt.Printf("%.0f\n", biaya)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/biaya/output.png)


#### Deskripsi
Program ini digunakan untuk menghitung perkiraan biaya kuliah pada tahun ke-n berdasarkan biaya awal sebesar Rp10.000.000 dan kenaikan sebesar 6% setiap tahun. Program menerima masukkan berupa tahun yang ingin dihitung, kemudian menggunakan perulangan untuk menaikkan biaya kuliah berdasarkan biaya tahun sebelumnya. Setiap perulangan mengalikan biaya dengan 1,06 dan menambah penghitung tahun hingga mencapai tahun yang ditentukan. Hasil akhir ditampilkan menggunakan fmt.Printf dengan format %.0f agar tidak menampilkan angka desimal. Program ini menerapkan konsep perulangan, operasi aritmatika, variabel bertipe float64, dan pemrosesan data secara berulang.

## Kesimpulan
Berdasarkan tugas praktikum yang telah dikerjakan, dapat disimpulkan bahwa perulangan digunakan untuk menjalankan instruksi secara berulang sesuai dengan kondisi yang ditentukan. Penggunaan perulangan for, percabangan if-else, variabel, dan operator perbandingan membantu program menyelesaikan berbagai permasalahan, seperti menentukan kelulusan, menghitung jumlah deret, mencari dua nilai terendah, serta menghitung kenaikan biaya kuliah. Dengan memahami konsep tersebut, penyusunan algoritma dan implementasinya dalam bahasa pemrograman Go menjadi lebih terstruktur dan efisien.