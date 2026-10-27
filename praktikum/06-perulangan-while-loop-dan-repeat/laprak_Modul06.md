# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan While</h1>
<p align="center">[Alfaro Steve Christian Hadiwiyata] - [109092600022]</p>

## Dasar Teori

### A. Pengertian Perulangan (Looping)
Perulangan atau looping merupakan struktur kontrol dalam algoritma pemrograman yang digunakan untuk menjalankan sekumpulan instruksi secara berulang selama kondisi tertentu terpenuhi. Perulangan membantu menyederhanakan penulisan program karena instruksi yang sama tidak perlu ditulis berulang kali. Dalam bahasa pemrograman Go, perulangan menggunakan kata kunci for, yang dapat diterapkan dengan tiga komponen maupun hanya menggunakan kondisi.

### B. Perulangan While dalam bahasa Go
Perulangan while merupakan struktur perulangan yang menjalankan instruksi selama kondisi bernilai benar (true). Bahasa pemrograman Go tidak menyediakan kata kunci while secara langsung. Oleh karena itu, perulangan while diterapkan menggunakan for yang hanya berisi kondisi. Kondisi diperiksa sebelum instruksi di dalam perulangan dijalankan. Jika kondisi bernilai benar, instruksi akan dijalankan. Sebaliknya, jika kondisi bernilai salah, perulangan akan berhenti. Oleh karena itu, diperlukan perubahan nilai variabel atau kondisi lain di dalam perulangan agar program dapat mencapai kondisi berhenti dan tidak mengalami perulangan tanpa akhir.

### C. Komponen Perulangan While
Perulangan while memiliki beberapa komponen penting, yaitu kondisi, variabel pengendali, dan instruksi yang diulang. Kondisi berfungsi menentukan apakah perulangan dilanjutkan atau dihentikan. Variabel pengendali digunakan untuk menyimpan nilai yang memengaruhi kondisi perulangan, sedangkan instruksi yang diulang berisi proses yang akan dijalankan selama kondisi masih terpenuhi.

### D. Perulangan Repeat-Until
Perulangan repeat-until merupakan struktur perulangan yang menjalankan instruksi terlebih dahulu, kemudian memeriksa kondisi berhenti. Dengan demikian, instruksi di dalam perulangan dijalankan setidaknya satu kali, meskipun kondisi berhenti sebenarnya sudah terpenuhi sejak awal.

Bahasa pemrograman Go tidak menyediakan kata kunci repeat-until secara langsung. Perilaku tersebut dapat diterapkan menggunakan for dengan kondisi penghentian yang diperiksa di dalam blok perulangan.

### E. Penerapan Perulangan While
Perulangan while dapat digunakan untuk menyelesaikan berbagai permasalahan, seperti mencari bilangan yang memenuhi batas tertentu, menghitung jumlah karakter dalam sebuah kalimat, dan mengulang proses input sampai kondisi tertentu tercapai. Dalam penerapannya, program perlu menentukan kondisi awal, proses yang diulang, serta kondisi berhenti.

Sebagai contoh, untuk mencari bilangan bulat terkecil yang pangkat tiganya lebih besar dari suatu batas, program dapat memulai pencarian dari bilangan 1. Selama hasil pangkat tiga belum melebihi batas, nilai bilangan akan terus ditambah. Setelah kondisi terpenuhi, program menampilkan bilangan yang ditemukan.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. terkecil.go

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
```
#### Deskripsi
Program ini digunakan untuk mencari bilangan bulat terkecil yang pangkat tiganya lebih besar dari batas yang dimasukkan pengguna. Variabel batas menyimpan nilai masukan, sedangkan n dimulai dari 1.

Perulangan for n*n*n <= batas berjalan selama pangkat tiga n masih kurang dari atau sama dengan batas. Setiap iterasi menambah n sebesar 1. Setelah kondisi tidak terpenuhi, program menampilkan nilai n sebagai hasil. Contohnya, jika batas adalah 12000, keluaran program adalah 23 karena 23^3 = 12167.

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}
```
#### Deskripsi
Program ini digunakan untuk mencari bilangan bulat terbesar yang kuadratnya lebih kecil dari batas masukan. Variabel batas menyimpan nilai masukan, sedangkan n dimulai dari 1.

Perulangan for n*n < batas terus berjalan selama kuadrat n masih lebih kecil dari batas. Nilai n bertambah satu pada setiap iterasi. Setelah kondisi tidak terpenuhi, program menampilkan n - 1 karena nilai n terakhir sudah tidak memenuhi syarat. Jika batas adalah 12000, hasilnya 109 karena 109^2 = 11881.

### 3. vokal.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```
#### Deskripsi
Program ini digunakan untuk menghitung jumlah huruf vokal (a, i, u, e, o) dalam sebuah kata atau kalimat, termasuk huruf kapital. Program menggunakan bufio.NewReader() dan ReadString('\n') untuk membaca satu baris teks.

Variabel i digunakan sebagai indeks karakter, sedangkan jumlahVokal menghitung jumlah huruf vokal. Perulangan berjalan selama i kurang dari panjang kalimat. Setiap karakter diubah menjadi huruf kecil, kemudian diperiksa menggunakan if. Jika karakter merupakan huruf vokal, jumlahVokal bertambah satu. Setelah semua karakter diperiksa, program menampilkan jumlah huruf vokal yang ditemukan.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. atm.go

```go
package main
import "fmt"
func main(){

    const passwordBenar = "1234"

    saldo := 10000000
    percobaan := 0
    login := false

    for percobaan < 3 && !login {
        var password string

        fmt.Print("Masukkan password ATM: ")
        fmt.Scan(&password)

        if password == passwordBenar {
            login = true
            fmt.Println("Login berhasil!")
        } else {
            percobaan++
            fmt.Println("Password salah. Silakan coba lagi.")
        }
    }

    if !login {
        fmt.Println("ATM terblokir. Program selesai.")
        return
    }

    pilihan := 0

    for pilihan != 3 {
        fmt.Println("Menu ATM:")
        fmt.Println("1. Lihat Saldo")
        fmt.Println("2. Tarik Uang")
        fmt.Println("3. Keluar")
        fmt.Print("Pilih: ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            fmt.Println("Saldo Anda: Rp", saldo)
        } else if pilihan == 2 {
            var penarikan int

            fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
            fmt.Scan(&penarikan)

            if penarikan <= 0 {
                fmt.Println("Jumlah penarikan tidak valid.")
            } else if penarikan <= saldo {
                saldo -= penarikan
                fmt.Println("Anda berhasil menarik uang sebesar Rp", penarikan)
                fmt.Println("Saldo tersisa: Rp", saldo)
            } else {
                fmt.Println("Saldo tidak mencukupi.")
            }
        } else if pilihan == 3 {
            fmt.Println("Keluar dari sistem. Terima kasih!")
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/alfsteve/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/unguided/atm/output.png)


#### Deskripsi
Program atm.go digunakan untuk membuat simulasi sistem ATM sederhana dengan fitur login, melihat saldo, menarik uang, dan keluar dari sistem. Program menerapkan perulangan while menggunakan for dengan kondisi untuk membatasi percobaan login dan menampilkan menu ATM secara berulang.

Pertama, program menetapkan konstanta passwordBenar dengan nilai "1234". Variabel saldo diinisialisasi dengan nilai 10000000 sebagai saldo awal pengguna. Variabel percobaan digunakan untuk menghitung jumlah kesalahan saat memasukkan password, sedangkan variabel login digunakan untuk menunjukkan apakah pengguna telah berhasil masuk ke sistem.

Perulangan for percobaan < 3 && !login digunakan untuk proses autentikasi. Perulangan berlangsung selama jumlah kesalahan kurang dari tiga dan pengguna belum berhasil login. Pada setiap iterasi, program meminta pengguna memasukkan password melalui fmt.Scan(&password). Percabangan if-else kemudian membandingkan password yang dimasukkan dengan password yang benar. Jika sesuai, variabel login diubah menjadi true dan program menampilkan pesan keberhasilan. Jika tidak sesuai, nilai percobaan bertambah satu dan program menampilkan pesan bahwa password salah.

Setelah perulangan login berakhir, program memeriksa nilai login. Jika pengguna belum berhasil masuk, program menampilkan pesan bahwa ATM terblokir, kemudian menghentikan fungsi main menggunakan return. Jika pengguna berhasil login, program melanjutkan ke menu ATM.

Variabel pilihan diinisialisasi dengan nilai 0. Perulangan for pilihan != 3 digunakan untuk menampilkan menu secara berulang sampai pengguna memilih pilihan ketiga. Di dalam perulangan, program menampilkan tiga pilihan, yaitu melihat saldo, menarik uang, dan keluar dari sistem.

Jika pengguna memilih angka 1, program menampilkan saldo yang tersimpan dalam variabel saldo. Jika pengguna memilih angka 2, program meminta jumlah uang yang akan ditarik dan menyimpannya dalam variabel penarikan. Percabangan memeriksa apakah jumlah penarikan kurang dari atau sama dengan nol. Jika demikian, program menampilkan pesan bahwa jumlah penarikan tidak valid. Jika jumlah penarikan positif dan tidak melebihi saldo, saldo dikurangi menggunakan saldo -= penarikan. Apabila jumlah penarikan melebihi saldo, program menampilkan pesan bahwa saldo tidak mencukupi.

Jika pengguna memilih angka 3, program menampilkan pesan keluar dari sistem. Karena kondisi perulangan menu adalah pilihan != 3, perulangan akan berhenti setelah pilihan tersebut dimasukkan. Sementara itu, jika pengguna memasukkan angka selain 1, 2, atau 3, program menampilkan pesan bahwa pilihan tidak valid dan kembali menampilkan menu.

## Kesimpulan
Berdasarkan praktikum Modul 06 tentang Perulangan While, dapat disimpulkan bahwa perulangan while dalam Go diterapkan menggunakan for yang hanya berisi kondisi. Perulangan ini digunakan untuk menjalankan instruksi selama kondisi masih terpenuhi.

Program terkecil.go dan terbesar.go menerapkan perulangan untuk mencari bilangan yang memenuhi batas tertentu, sedangkan vokal.go menggunakannya untuk menghitung huruf vokal dalam kalimat. Program atm.go menerapkan perulangan untuk membatasi percobaan login dan menampilkan menu ATM hingga pengguna memilih keluar.

Praktikum ini memberikan pemahaman tentang penggunaan kondisi, perubahan variabel, dan penghentian perulangan, serta penerapannya bersama percabangan untuk menyelesaikan permasalahan menggunakan bahasa pemrograman Go.

## Referensi
1. The Go Authors. (n.d.). A Tour of Go — Basics. https://go.dev/tour/basics
2. The Go Authors. (n.d.). A Tour of Go — Flow Control Statements: for, if, else. https://go.dev/tour/flowcontrol
3. The Go Authors. (n.d.). Package fmt Documentation. https://pkg.go.dev/fmt
4. The Go Authors. (n.d.). Package bufio Documentation. https://pkg.go.dev/bufio
5. The Go Authors. (n.d.). Package strings Documentation. https://pkg.go.dev/strings