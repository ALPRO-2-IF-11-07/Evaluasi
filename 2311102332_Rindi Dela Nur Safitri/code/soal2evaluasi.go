package main

import (
    "fmt"
)

func hitungBiayaMinuman(banyakOrang string, jenisMinuman int, sisa bool) int {
    var tarifPerMinuman float64
    if jenisMinuman == "Dasar" {
        tarifPerMinuman = 15000
    } else if jenisMinuman == "Menengah" {
        tarifPerMinuman = 75000
	} else if jenisMinuman == "Besar" {
		tarifPerMinuman = 150000
    } else {
        fmt.Println("Jenis Minuman tidak valid.")
        return 0
    }
	
    totalbanyakOrang := float64 (banyakOrang)
    if banyakOrang >= 5 {
        totalOrang += 5
    } else if banyakOrang >= 10 {
        totalOrang += 10
    }

    totalBiaya := totalbanyakOrang * tarifPerMinuman

    if totalJenis == "Tidak Habis" {
        tarifBiaya *= 0.8
    }

    return int(totalBiaya)
}

func main() {
    var banyakOrang string
    var jenisMinuman int
	var sisa bool

	fmt.Print("Masukkan jumlah kelompok: ")
    fmt.Scan(&N)
    fmt.Print("Masukkan jumlah jenis minuman: ")
    fmt.Scan(&jenisMinuman)
    fmt.Print("Masukkan banyak orang: ")
    fmt.Scan(&banyakOrang)
    fmt.Print("Apakah sisa? (true/false): ")
    fmt.Scan(&sisa)

    biaya := hitungBiayaMinuman(jenisMinuman, banyakOrang, sisa)

    fmt.Printf("Biaya Minuman: Rp %d\n", biaya)
}

