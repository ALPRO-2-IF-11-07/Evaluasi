// no 1
// package main

// import (
// 	"fmt"
// )

// func hitungBiayaTiket(jumlahTiket int, jenisKursi string, isMember bool) int {
// 	var hargaPerTiket int

// 	if jenisKursi == "biasa" {
// 		if isMember {
// 			hargaPerTiket = 40000
// 		} else {
// 			hargaPerTiket = 50000 
// 		}
// 	} else if jenisKursi == "VIP" {
// 		if isMember {
// 			hargaPerTiket = 60000
// 		} else {
// 			hargaPerTiket = 70000
// 		}
// 	} else {
// 		fmt.Println("Jenis kursi tidak valid.")
// 		return 0
// 	}

// 	totalBiaya := jumlahTiket * hargaPerTiket

// 	if jumlahTiket > 2 {
// 		totalBiaya = int(float64(totalBiaya) * 0.85) 
// 	}

// 	return totalBiaya
// }

// func main() {
// 	var jumlahTiket int
// 	var jenisKursi string
// 	var isMember bool

// 	fmt.Print("Masukkan jumlah tiket: ")
// 	fmt.Scan(&jumlahTiket)
// 	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
// 	fmt.Scan(&jenisKursi)
// 	fmt.Print("Apakah Anda member? (true/false): ")
// 	fmt.Scan(&isMember)

// 	biaya := hitungBiayaTiket(jumlahTiket, jenisKursi, isMember)

// 	fmt.Printf("Total biaya tiket: Rp %d\n", biaya)
// }


// no2
// package main

// import (
// 	"fmt"
// )

// func isBungaSempurna(jumlahKelopak int) bool {
// 	var totalKelopak int

// 	for i := 1; i < jumlahKelopak; i++ {
// 		if jumlahKelopak%i == 0 { 
// 			totalKelopak += i
// 		}
// 	}

// 	return totalKelopak == jumlahKelopak
// }

// func findBungaSempurna(a, b int) []int {
// 	var bungaSempurna []int

// 	for i := a; i <= b; i++ {
// 		if isBungaSempurna(i) {
// 			bungaSempurna = append(bungaSempurna, i)
// 		}
// 	}

// 	return bungaSempurna
// }

// func main() {
// 	var a, b int

// 	fmt.Print("Masukkan jumlah kelopak (a): ")
// 	fmt.Scan(&a)
// 	fmt.Print("Masukkan jumlah kelopak (b): ")
// 	fmt.Scan(&b)

// 	if a <= 0 || b <= 0 || a > b {
// 		fmt.Println("Input tidak valid. Pastikan a dan b adalah bilangan bulat positif dan a <= b.")
// 		return
// 	}

// 	bungaSempurna := findBungaSempurna(a, b)

// 	if len(bungaSempurna) > 0 {
// 		fmt.Printf("Bunga sempurna antara %d dan %d: ", a, b)
// 		for _, bunga := range bungaSempurna {
// 			fmt.Printf("%d ", bunga)
// 		}
// 		fmt.Println()
// 	} else {
// 		fmt.Printf("Tidak ada bunga sempurna dengan jumlah kelopak antara %d dan %d.\n", a, b)
// 	}
// }

// no 3
package main

import (
	"fmt"
)

func hitungSesiPelatihan(hari, p, q int) int {

	if hari > 365 {
		return 0
	}

	if hari%p == 0 && hari%q != 0 {

		return 1 + hitungSesiPelatihan(hari+1, p, q)
	}

	return hitungSesiPelatihan(hari+1, p, q)
}

func main() {
	var p, q int

	fmt.Print("Masukkan kelipatan p (hari sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan kelipatan q (hari tidak ada sesi pelatihan): ")
	fmt.Scan(&q)

	if p <= 0 || q <= 0 {
		fmt.Println("Kelipatan p dan q harus bilangan bulat positif.")
		return
	}

	jumlahSesi := hitungSesiPelatihan(1, p, q)

	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", jumlahSesi)
}
