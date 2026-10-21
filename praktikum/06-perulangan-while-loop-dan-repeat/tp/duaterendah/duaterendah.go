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