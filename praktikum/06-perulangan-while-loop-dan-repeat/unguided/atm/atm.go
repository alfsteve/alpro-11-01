package main
import "fmt"
func main(){

    const passwordBenar = "1234"

    saldo := 10000000
    percobaan := 0
    login := false

    for percobaan < 3 && !login {
        var password string

        fmt.Print("Masukkan password ATM: ")
        fmt.Scan(&password)

        if password == passwordBenar {
            login = true
            fmt.Println("Login berhasil!")
        } else {
            percobaan++
            fmt.Println("Password salah. Silakan coba lagi.")
        }
    }

    if !login {
        fmt.Println("ATM terblokir. Program selesai.")
        return
    }

    pilihan := 0

    for pilihan != 3 {
        fmt.Println("Menu ATM:")
        fmt.Println("1. Lihat Saldo")
        fmt.Println("2. Tarik Uang")
        fmt.Println("3. Keluar")
        fmt.Print("Pilih: ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            fmt.Println("Saldo Anda: Rp", saldo)
        } else if pilihan == 2 {
            var penarikan int

            fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
            fmt.Scan(&penarikan)

            if penarikan <= 0 {
                fmt.Println("Jumlah penarikan tidak valid.")
            } else if penarikan <= saldo {
                saldo -= penarikan
                fmt.Println("Anda berhasil menarik uang sebesar Rp", penarikan)
                fmt.Println("Saldo tersisa: Rp", saldo)
            } else {
                fmt.Println("Saldo tidak mencukupi.")
            }
        } else if pilihan == 3 {
            fmt.Println("Keluar dari sistem. Terima kasih!")
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    }
}