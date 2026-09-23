package main

	import "fmt"

		func main() {

		var a, b int
	
		//Membaca input

		fmt.Print("Masukkan nilai a:")
			fmt.Scan(&a)
		
		fmt.Print("Masukkan nilai b:")
			fmt.Scan(&b)
		
		//Menukar nilai a dan b

			a, b = b, a

		//Menampilkan output

			fmt.Println(a)
			fmt.Println(b)

}