package main

import "fmt"

func tarifTiket(jenisKursi string, jumlahTiket float64, rilMemberOrFek bool) int {
	var tarifTiket, raisa float64

	if rilMemberOrFek {
		if jenisKursi == "biasa" {
			tarifTiket = 40000
		} else if jenisKursi == "VIP" {
			tarifTiket = 60000
		} else {
			fmt.Println("Tidak Valid")
			return 0
		}
	} else if rilMemberOrFek == false {
		if jenisKursi == "biasa" {
			tarifTiket = 50000
		} else if jenisKursi == "VIP" {
			tarifTiket = 70000
		} else {
			fmt.Println("Tidak Valid")
			return 0
		}
	}

	raisa = tarifTiket * jumlahTiket

	if jumlahTiket > 2 {
		raisa *= 0.85
	}

	return int(raisa)
}

func main() {
	var jenisKursi string
	var jumlahTiket int
	var rilMemberOrFek bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&jumlahTiket)
	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scan(&jenisKursi)
	fmt.Print("Apakah anda member? (true/false): ")
	fmt.Scan(&rilMemberOrFek)

	biaya := tarifTiket(jenisKursi, float64(jumlahTiket), rilMemberOrFek)

	fmt.Printf("Total biaya: %d\n", biaya)
}
