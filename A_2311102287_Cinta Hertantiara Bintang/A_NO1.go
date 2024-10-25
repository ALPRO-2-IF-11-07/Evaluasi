package main

import (
	"fmt"
)

func hitungBiayaPembelian(tiket int, jenisKursi string, isMember bool) int {
	var tarifPerKursi float64

	if jenisKursi == "biasa" {
		if isMember {
			tarifPerKursi = 40000
		} else {
			tarifPerKursi = 50000
		}
	} else if jenisKursi == "VIP" {
		if isMember {
			tarifPerKursi = 60000
		} else {
			tarifPerKursi = 70000
		}
	} else {
		fmt.Println("Jenis Kursi tidak valid.")
		return 0
	}

	totalBiaya := float64(tiket) * tarifPerKursi

	if tiket > 2 {
		totalBiaya *= 0.85
	}

	return int(totalBiaya)
}

func main() {
	var tiket int
	var jenisKursi string
	var isMember bool

	fmt.Print("Masukkan jumlah tiket : ")
	fmt.Scan(&tiket)
	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scan(&jenisKursi)
	fmt.Print("Apakah Anda Member? (true/false) : ")
	fmt.Scan(&isMember)

	biaya := hitungBiayaPembelian(tiket, jenisKursi, isMember)

	fmt.Printf("Biaya Pembelian Tiket : Rp. %d\n", biaya)
}
