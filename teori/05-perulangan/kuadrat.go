package main

import "fmt"

func main() {
    var angka int
    fmt.Scan(&angka)

    total := 0

    for angka > 0 {
        digit := angka % 10
        kuadrat := digit * digit

        fmt.Print(kuadrat, " ")
        total = total + kuadrat

        angka = angka / 10
    }

    fmt.Println()
    fmt.Println(total)
}