package main

import (
	"fmt"
)

func hitungBiayaPembelian(jumlahTiket int, jenisKursi string, isMember bool) int {
	var hargaTiket int

	if jenisKursi == "biasa" {
		if isMember {
			hargaTiket = 40000
		} else {
			hargaTiket = 50000
		}
	} else if jenisKursi == "vip" {
		if isMember {
			hargaTiket = 60000
		} else {
			hargaTiket = 70000
		}
	} else {
		fmt.Println("Jenis kursi tidak valid")
		return 0
	}

	//sebelum diskon
	totalBiaya := jumlahTiket * hargaTiket

	// Berikan diskon 15% jika membeli lebih dari 2 tiket
	if jumlahTiket > 2 {
		totalBiaya = int(float64(totalBiaya) * 0.85)
	}

	return totalBiaya
}

func main() {
	var jumlahTiket int
	var jenisKursi string
	var isMember bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&jumlahTiket)
	fmt.Print("Masukkan jenis kursi (biasa/vip): ")
	fmt.Scan(&jenisKursi)
	fmt.Print("Apakah anda member? (true/false): ")
	fmt.Scan(&isMember)

	total := hitungBiayaPembelian(jumlahTiket, jenisKursi, isMember)

	fmt.Printf("Pembelian Tiket: Rp %d\n", total)
}
