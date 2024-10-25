package main

import "fmt"

func bungaSempurna(n int) bool {
	raisa := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			raisa += i
		}
	}
	return raisa == n
}

func main() {
	var a, b int

	fmt.Print("Masukkan jumlah kelopak (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah kelopak (b): ")
	fmt.Scan(&b)

	fmt.Printf("Bunga Sempurna antara %d dan %d:", a, b)
	for i := a; i <= b; i++ {
		if bungaSempurna(i) {
			fmt.Print(" ", i)
		}
	}
}
