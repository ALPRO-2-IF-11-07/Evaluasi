package main

import (
	"fmt"
)

func hitungSesi(p, q, hari int) int {
	if hari == 0 {
		return 0
	}
	if hari%p == 0 && hari%q != 0 {
		return 1 + hitungSesi(p, q, hari-1)
	}
	return hitungSesi(p, q, hari-1)
}

func main() {
	var p, q int

	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&q)

	totalSesi := hitungSesi(p, q, 365)

	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", totalSesi)
}
