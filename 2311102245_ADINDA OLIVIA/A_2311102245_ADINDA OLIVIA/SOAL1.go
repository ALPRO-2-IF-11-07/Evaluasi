package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung hasil bagi dan sisa pembagian
func bagiRekursif(n int, m int) (int, int) {
	// Basis kasus: jika pembilang lebih kecil dari penyebut, hasil bagi adalah 0 dan sisanya adalah n
	if n < m {
		return 0, n
	}
	// Rekurens: kurangi n dengan m dan tambahkan 1 ke hasil bagi
	hasilBagi, sisa := bagiRekursif(n-m, m)
	return hasilBagi + 1, sisa
}

func main() {
	// Input dari pengguna
	var n, m int
	fmt.Print("Masukkan bilangan pembilang (n): ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan penyebut (m): ")
	fmt.Scan(&m)

	// Periksa jika masukan adalah bilangan bulat positif
	if n <= 0 || m <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	// Panggil fungsi rekursif untuk menghitung hasil bagi dan sisa
	hasilBagi, sisa := bagiRekursif(n, m)

	// Tampilkan hasil pembagian dan sisa
	fmt.Printf("Hasil dari pembagian %d dengan %d adalah: %d\n", n, m, hasilBagi)
	fmt.Printf("Sisa dari pembagian adalah: %d\n", sisa)
}
