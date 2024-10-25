package main

import (
	"fmt"
)

// Fungsi untuk menghitung biaya berdasarkan jumlah minuman, jumlah orang, dan sisa minuman
func calculateCost(jumlahMinuman, jumlahOrang int, sisa bool) int {
	var totalCost int

	// Hitung biaya berdasarkan jumlah minuman
	if jumlahMinuman <= 5 {
		totalCost = jumlahMinuman * 15000
	} else if jumlahMinuman <= 10 {
		totalCost = 75000 + (jumlahMinuman-5)*10000
	} else {
		totalCost = 150000
	}

	// Tambahkan biaya tambahan jika ada sisa minuman
	if sisa {
		biayaTambahan := int(float64(totalCost) * 0.20 * float64(jumlahOrang))
		totalCost += biayaTambahan
	}

	return totalCost
}

func main() {
	var M int
	fmt.Print("Masukkan jumlah kelompok: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var jumlahMinuman, jumlahOrang int
		var sisa bool

		fmt.Printf("Masukkan jumlah jenis minuman, banyak orang, dan apakah sisa (true/false) untuk kelompok %d: ", i)
		fmt.Scan(&jumlahMinuman, &jumlahOrang, &sisa)

		totalCost := calculateCost(jumlahMinuman, jumlahOrang, sisa)
		fmt.Printf("Total biaya untuk kelompok %d: Rp %d\n", i, totalCost)
	}
}