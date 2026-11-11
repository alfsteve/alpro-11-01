package main

import "fmt"

func main() {
	var n, x, min int

	fmt.Scan(&n)

	fmt.Scan(&min)

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x < min {
			min = x
		}
	}

	fmt.Println(min)
}