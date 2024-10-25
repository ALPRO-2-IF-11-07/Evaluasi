package main

import (
	"fmt"
)

func BungaSempuran(n int) bool {
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

	fmt.Print("Masukkan jumlah kelopak (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah kelopak (b): ")
	fmt.Scan(&b)

	fmt.Printf("Bunga Sempurna antara %d dan %d: ", a, b)

	found := false
	for i := a; i <= b; i++ {
		if BungaSempuran(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}

	if !found {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}
