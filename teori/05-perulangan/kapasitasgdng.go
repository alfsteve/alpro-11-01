package main

import "fmt"

func main() {
    var kapasitas int
    fmt.Scan(&kapasitas)

    total := 0

    for i := 1; i <= 5; i++ {
        var barang int
        fmt.Scan(&barang)

        total = total + barang

        fmt.Println(total >= kapasitas)
    }
}