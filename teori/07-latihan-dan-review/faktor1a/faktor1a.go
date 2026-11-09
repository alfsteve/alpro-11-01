package main
import "fmt"
func main() {
	var N, M int
	fmt.Println("Masukkan nilai N dan M:")
	fmt.Scan(&N, &M)

	batas := N
	if M < N {
		batas = M
	}

	for i := 1; i <= batas; i++ {
		if N%i == 0 && M%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
