package main

import "fmt"

func main() {
	var jenisKursi string
	var jumlahTiket int
	var TrueorFalse bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&jumlahTiket)
	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scan(&jenisKursi)
	fmt.Print("Apakah anda member? (true/false): ")
	fmt.Scan(&TrueorFalse)

	biaya := tarifTiket(jenisKursi, (jumlahTiket), TrueorFalse)

	fmt.Printf("Total biaya: %d\n", biaya)
}

func tarifTiket(jenisKursi string, jumlahTiket, TrueorFalse bool) int {
	var TrueorFalse, 

	if TrueorFalse {
		if jenisKursi == "biasa" {
			tarifTiket = 40000
		} else if jenisKursi == "VIP" {
			tarifTiket = 60000
		} else {
			fmt.Println("Tidak Valid")
			return 0
		}
	} else if TrueorFalse == false {
		if jenisKursi == "biasa" {
			tarifTiket = 50000
		} else if jenisKursi == "VIP" {
			tarifTiket = 70000
		} else {
			fmt.Println("Tidak Valid")
			return 0
		}
	}


	return int
}


