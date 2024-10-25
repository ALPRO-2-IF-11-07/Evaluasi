package main

import "fmt"

func hitungtariftiket(jumlahtiket int, jenis string, ismember bool) int {
	var tarif float64
	if jenis == "biasa" {
		tarif = 40000
	} else if jenis == "mewah" {
		tarif = 60000
	} else {
		fmt.Println("tidak valid")
		return 0
	}

	totaltiket := float64(jumlahtiket)
	totalbiaya := totaltiket * tarif

	if totaljam > 1 {
		totalbiaya *= 0.15
	}

	// Diskon untuk member
	if ismember {
		totaltiket *= 0.85
	}

	return int(totaltiket)
}

func main() {
	var jumlahtiket int
	var jenis string
	var ismember bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&jumlahtiket)
	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scan(&jenisn)
	fmt.Print("Apakah Anda member (true/false): ")
	fmt.Scan(&ismember)

	biaya := hitungtariftiket(jumlahtiket, jeniskendaraan, ismember)

	fmt.Printf("Biaya parkir: Rp%d\n", biaya)
}
