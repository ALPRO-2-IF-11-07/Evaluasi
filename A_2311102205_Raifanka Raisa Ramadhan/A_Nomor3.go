package main

import "fmt"

func hitungSesi(p, q, raisa int) int {
	if raisa == 0 {
		return 0
	}

	if raisa%p == 0 && raisa%q != 0 {
		return 1 + hitungSesi(p, q, raisa-1)
	}

	return hitungSesi(p, q, raisa-1)
}

func main() {
	var p, q int

	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&q)

	jumlahSesi := hitungSesi(p, q, 365)

	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", jumlahSesi)
}
