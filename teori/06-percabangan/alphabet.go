package main
import "fmt"
func main(){
	var karakter int
	fmt.Scan("%c", &karakter)

	if karakter >= 'A' && karakter <= 'Z' || karakter >= 'a' && karakter <= 'z' {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}
