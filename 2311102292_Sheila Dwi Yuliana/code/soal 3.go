package main

import (
	"fmt"
)

func hitungSesi(hari, p, q, total int) int {

	if hari > 365 {
		return total
	}

	if hari%p == 0 && hari%q != 0 {
		return hitungSesi(hari+1, p, q, total+1)
	}

	return hitungSesi(hari+1, p, q, total)
}

func main() {
	var p, q int

	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&q)

	if p <= 0 || q <= 0 {
		fmt.Println("Nilai p dan q harus bilangan bulat positif.")
		return
	}

	totalSesi := hitungSesi(1, p, q, 0)

	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", totalSesi)
}
