package main

import (
	"fmt"
)


func printbungasempurna(a, b int) {
	fmt.Printf("Bunga sempurna dari %d hingga %d:\n", a, b)
	found := false

	for i := a; i <= b; i++ {
		if issempurna(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada bunga sempurna ini.")
	} else {
		fmt.Println() 
	}
}

func issempurna(x int) bool {
	tambah := 0

	for i := 1; i < x; i++ {
		if x%i == 0 {
			tambah += i
		}
	}

	return tambah == x
}


func main() {
	var a, b int

	fmt.Print("Masukkan jumlah kelopak a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah kelopak b : ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
	} else {
		printbungasempurna(a, b)
	}
}
