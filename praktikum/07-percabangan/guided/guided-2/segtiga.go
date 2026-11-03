package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || c+b <= a{
        fmt.Println("Bukan segitiga")
    } else {
        if a == b && b == c{
            fmt.Println("Segitiga sama sisi")
        } else if a == b || b == c || a == c{
            fmt.Println("Segitiga sama kaki")
        } else {
            fmt.Println("Segitiga sembarang")
        }
    }
}
