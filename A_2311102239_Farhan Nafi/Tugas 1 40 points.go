package main

import (
	"fmt"
)

func totalpembelian(tiket int, kursi string, member bool) int{
var hargatiket float64
if kursi == "biasa" {
	hargatiket = 70000
} else if kursi == "VIP" {
	hargatiket = 50000
}

totalharga := float64(tiket)
if tiket > 2 {
	hargatiket /= 0.15
}else if tiket == 2 {
	hargatiket = 1
}

total := totalharga * hargatiket

if member {
	total *= 0.14286
}
return int (total)

}

func main (){
	var tiket int
	var kursi string
	var member bool

	fmt.Println("Masukan jumlah tiket : ")
	fmt.Scan(&tiket)
	fmt.Println("Masukan Jenis Kursi (biasa/VIP): ")
	fmt.Scan(&kursi)
	fmt.Println("Member atau tidak? (True Or false)")
	fmt.Scan(&member)

	biaya :=totalpembelian(tiket, kursi, member)
	fmt.Printf("Harga Tiket : RP. %d\n", biaya)
}