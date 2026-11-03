
package main

import "fmt"

func main() {
    var a, b, c int

    fmt.Print("Masukkan tiga bilangan: ")
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println("Bilangan terbesar:", a)
    } else if b >= a && b >= c {
        fmt.Println("Bilangan terbesar:", b)
    } else {
        fmt.Println("Bilangan terbesar:", c)
    }
}
