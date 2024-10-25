package main

import (
	"fmt"
	"strconv"
)
//HANAH NUR AZIZAH 2311102312
func isEven(n int) bool {
	return n%2 == 0
}
func isOdd(n int) bool {
	return n%2 != 0
}
func hadiahTiket(tiket string) string {
	allEven := true
	allOdd := true
	for _, char := range tiket {
		digit, _ := strconv.Atoi(string(char))
		if isEven(digit) {
			allOdd = false
		} else {
			allEven = false
		}
	}
	if allEven {
		return "Hadiah Utama"
	} else if allOdd {
		return "Hadiah Sembako"
	} else {
		return "Hadiah Konsol"
	}
}
func main() {
	var N int
	fmt.Print("Masukkan jumlah peserta acara: ")
	fmt.Scan(&N)
	hadiahUtamaCount := 0
	hadiahSembakoCount := 0
	hadiahKonsolCount := 0
	for i := 1; i <= N; i++ {
		var tiket string
		fmt.Printf("Masukkan nomor tiket peserta %d: ", i)
		fmt.Scan(&tiket)
		hadiah := hadiahTiket(tiket)
		fmt.Println(hadiah)
		switch hadiah {
		case "Hadiah Utama":
			hadiahUtamaCount++
		case "Hadiah Sembako":
			hadiahSembakoCount++
		case "Hadiah Konsol":
			hadiahKonsolCount++
		}
	}
	fmt.Println("\nJumlah:")
	fmt.Printf("Hadiah Utama: %d\n", hadiahUtamaCount)
	fmt.Printf("Hadiah Sembako: %d\n", hadiahSembakoCount)
	fmt.Printf("Hadiah Konsol: %d\n", hadiahKonsolCount)
}
