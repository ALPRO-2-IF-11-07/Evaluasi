package main

import (
	"fmt"
)

func hitunghari(day, x, y int) int {
	if day > 365 {
		return 0
	}

	if day%x == 0 && day%y != 0 {
		// Hitung pertemuan hari ini dan lanjutkan ke hari berikutnya
		return 1 + hitunghari(day+1, x, y)
	}

	// Jika tidak bertemu, lanjutkan ke hari berikutnya
	return hitunghari(day+1, x, y)
}

func main() {
	var x, y int

	// Input dari pengguna
	fmt.Print("Masukkan nilai p (kelipatan): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&y)

	jumlahPertemuan := hitunghari(1, x, y)

	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", jumlahPertemuan)
}
