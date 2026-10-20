
package main

import "fmt"

func main() {
    var mahasiswa int
    fmt.Scan(&mahasiswa)

    sertif := 0

    for i := 1; i <= mahasiswa; i++ {
        total := 0

        for j := 1; j <= 6; j++ {
            var nilai int
            fmt.Scan(&nilai)
            total += nilai
        }

        ratarata := float64(total) / 6.0

        if ratarata >= 3.33 {
            fmt.Println("Ya")
            sertif++
        } else {
            fmt.Println("Tidak")
        }
    }

    fmt.Println(sertif)
}