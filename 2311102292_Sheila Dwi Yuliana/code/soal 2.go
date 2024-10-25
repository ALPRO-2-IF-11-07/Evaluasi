package main

import (
	"fmt"
)

func bungaSempurna(n int) bool {
	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}

func main() {
	var a, b int

	fmt.Print("Masukkan rentang awal (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan rentang akhir (b): ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Rentang tidak valid. Pastikan a <= b.")
		return
	}

	fmt.Printf("Bunga sempurna antara %d hingga %d:\n", a, b)

	adaBungaSempurna := false
	for i := a; i <= b; i++ {
		if bungaSempurna(i) {
			fmt.Println(i)
			adaBungaSempurna = true
		}
	}

	if !adaBungaSempurna {
		fmt.Println("Tidak ada bunga sempurna ditemukan dalam rentang tersebut.")
	}
}
