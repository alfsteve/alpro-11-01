package main

import "fmt"

func main() {
	var L, R int
	var jumlah int

	fmt.Scan(&L, &R)

	for i := L; i <= R; i++ {

		x := i
		sum := 0
		valid := true

		for x > 0 {
			depan := (x / 10) % 10
			belakang := x % 10

			if depan > belakang {
				valid = false
			}

			sum += belakang
			x = x / 10
		}

		prima := true

		if sum < 2 {
			prima = false
		} else {
			for j := 2; j < sum; j++ {
				if sum%j == 0 {
					prima = false
				}
			}
		}

		if valid && prima {
			jumlah++
		}
	}

	fmt.Println(jumlah)
}