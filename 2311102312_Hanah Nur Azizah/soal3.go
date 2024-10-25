package main

import (
	"fmt"
)
//HANAH NUR AZIZAH 2311102312
func divide(n, m int) (int, int) {
	if n < m {
		return 0, n
	}
	quotient, remainder := divide(n-m, m)
	return quotient + 1, remainder
}
func main() {
	var n, m int
	fmt.Print("Masukkan bilangan pembilang (n): ")
	fmt.Scanln(&n)
	fmt.Print("Masukkan bilangan penyebut (m): ")
	fmt.Scanln(&m)
	if m == 0 {
		fmt.Println("Penyebut tidak boleh nol!")
		return
	}
	quotient, remainder := divide(n, m)
	fmt.Printf("Hasil dari pembagian %d dengan %d adalah: %d\n", n, m, quotient)
	fmt.Printf("Sisa dari pembagian adalah: %d\n", remainder)
}
