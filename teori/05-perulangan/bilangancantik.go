package main

import "fmt"

func main() {
    var angka int
    fmt.Scan(&angka)

    sebelumnya := 10
    indah := true

    for angka > 0 {
        digit := angka % 10

        if digit > sebelumnya {
            indah = false
        }

        sebelumnya = digit
        angka = angka / 10
    }

    fmt.Println(indah)
}
