package main
import "fmt"
func main(){

	var x, y, z int
	
	fmt.Scan(&x, &y, &z)
	temp := x
	z = y
	y = temp

	fmt.Println(x, y, z)
}
