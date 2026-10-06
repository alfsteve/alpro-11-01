package main

import "fmt"

func main() {

    var n, i, bilangan, jumlah, d1, d4 int
    // Membuat 6 variabel bertipe int (bilangan bulat):
    // n = jumlah bilangan yang akan dimasukkan
    // i = penghitung perulangan
    // bilangan = menyimpan bilangan yang diinput
    // jumlah = menyimpan hasil penjumlahan
    // d1 = menyimpan digit pertama
    // d4 = menyimpan digit terakhir

    fmt.Scan(&n)
    // Membaca input dari user dan menyimpannya ke variabel n

    jumlah = 0
    // Mengatur nilai awal jumlah menjadi 0

    for i = 1; i <= n; i++ {
        // Melakukan perulangan dari i = 1 sampai i <= n
        // i++ artinya nilai i bertambah 1 setiap perulangan

        fmt.Scan(&bilangan)
        // Membaca bilangan dari user dan menyimpannya ke variabel bilangan

        d1 = bilangan / 1000
        // Mengambil digit pertama dari bilangan 4 digit
        // Contoh: 1234 / 1000 = 1

        d4 = bilangan % 10
        // Mengambil digit terakhir dari bilangan
        // % adalah modulo (sisa pembagian)
        // Contoh: 1234 % 10 = 4

        jumlah += d1 + d4
        // Menambahkan digit pertama dan digit terakhir ke jumlah
        // Sama dengan: jumlah = jumlah + d1 + d4
    }

    fmt.Println(jumlah)
    // Menampilkan hasil akhir jumlah ke layar
}