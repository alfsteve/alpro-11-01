package main
import "fmt"
func main() {
	
    var kendaraan string
    var jam, tarif, denda, total int

    fmt.Print("Masukkan jenis kendaraan (Mobil/Motor): ")
    fmt.Scan(&kendaraan)

    fmt.Print("Masukkan lama parkir (jam): ")
    fmt.Scan(&jam)

    if jam <= 0 {
        fmt.Println("Lama parkir tidak valid")
        return
    }

    if kendaraan == "Mobil" {
        tarif = 5000
        denda = 50000

        if jam > 1 {
            total = tarif + (jam-1)*3000
        } else {
            total = tarif
        }

    } else if kendaraan == "Motor" {
        tarif = 3000
        denda = 30000

        if jam > 1 {
            total = tarif + (jam-1)*2000
        } else {
            total = tarif
        }

    } else {
        fmt.Println("Jenis kendaraan tidak valid")
        return
    }

    if jam > 24 {
        total += denda
    }

    fmt.Println("Total tarif parkir: Rp", total)
}
