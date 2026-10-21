package main

import "fmt"

func main() {
    var tepung, telur int
    var tepungPerDonat, telurPerDonat int

    fmt.Scan(&tepung, &telur)
    fmt.Scan(&tepungPerDonat, &telurPerDonat)

    jumlahDonat := 0

    for tepung >= tepungPerDonat &&
        telur >= telurPerDonat {

        tepung = tepung - tepungPerDonat
        telur = telur - telurPerDonat

        jumlahDonat++
    }

    fmt.Println(jumlahDonat)
}