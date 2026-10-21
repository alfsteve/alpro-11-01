package main

import "fmt"

func main() {
    var poin int
    total := 0

    fmt.Scan(&poin)

    for poin != 0 {
        total = total + poin
        fmt.Scan(&poin)
    }

    fmt.Println(total)
}