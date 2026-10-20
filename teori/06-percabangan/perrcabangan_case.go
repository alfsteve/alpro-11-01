package main
import "fmt"
func main(){

	var Rambu string
	fmt.Scan(&Rambu)

	if Rambu == "Merah"{
		fmt.Println("STOP!")
	} else if Rambu == "Kuning"{
		fmt.Println("Bersiaplah / Hati-hati")
	} else if Rambu == "Hijau"{
		fmt.Println("Gas Boskyuhhh")
	} else {
		fmt.Println("Perintah ga ada di list mas")
	}
}