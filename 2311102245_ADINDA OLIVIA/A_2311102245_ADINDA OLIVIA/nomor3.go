package main

import (
	"fmt"
)

func potongBilangan(bilangan int) (int, int, int) {
	temp := bilangan
	panjang := 0
	for temp > 0 {
		temp /= 10
		panjang++
	}

	var bagian1, bagian2, bagian3 int
	var pembagi1, pembagi2 int

	if panjang%2 == 0 {
		pembagi1 = 100000
		pembagi2 = 1000

		bagian1 = bilangan / pembagi1
		bagian2 = (bilangan % pembagi1) / pembagi2
		bagian3 = bilangan % pembagi2

	} else {
		pembagi1 = 10000
		pembagi2 = 100

		bagian1 = bilangan / pembagi1
		bagian2 = (bilangan % pembagi1) / pembagi2
		bagian3 = bilangan % pembagi2
	}
	return bagian1, bagian2, bagian3
}

func hitungRataRata(bag1, bag2, bag3 int) float64 {
	return float64(bag1+bag2+bag3) / 3
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif lebih dari 100: ")
	fmt.Scan(&bilangan)

	if bilangan <= 100 {
		fmt.Println("bilangan harus lebih dari 100")
		return
	}

	bagian1, bagian2, bagian3 := potongBilangan(bilangan)
	fmt.Printf("Hasil pemotongan: %d %d %d\n", bagian1, bagian2, bagian3)
	rataRata := hitungRataRata(bagian1, bagian2, bagian3)

	fmt.Printf("rata-ratanya adalah: %.2f\n", rataRata)
}
