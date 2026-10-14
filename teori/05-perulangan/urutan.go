package main

import "fmt"
func main() {
	var n, iterasi int
	fmt.Scan(&n)
	iterasi = 1
	for {
		fmt.Print(iterasi, " ")
		iterasi++
		if iterasi > n {
			break
		}
	}
}