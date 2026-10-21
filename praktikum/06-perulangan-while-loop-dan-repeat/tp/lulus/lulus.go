package main
import "fmt"
func main(){
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	for nilai != -1{
	if nilai >= 60{
		fmt.Println("Anda Lulus Ujian")
	} else {
		fmt.Println("Anda tidak Lulus")
	}
		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&nilai)
	}
	fmt.Println("Program selesai.")
}
