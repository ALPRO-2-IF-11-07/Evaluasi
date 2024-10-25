package main

import (
	"fmt"
	"strconv"
)

// Fungsi untuk mengecek apakah tiket valid
func isValidTicket(ticket string) bool {
	length := len(ticket)

	// Cek panjang tiket, harus 6 atau 8
	if length != 6 && length != 8 {
		return false
	}

	// Menghitung jumlah dari dua digit pertama dan dua digit terakhir
	firstDigit1, err1 := strconv.Atoi(string(ticket[0]))
	firstDigit2, err2 := strconv.Atoi(string(ticket[1]))
	lastDigit1, err3 := strconv.Atoi(string(ticket[length-2]))
	lastDigit2, err4 := strconv.Atoi(string(ticket[length-1]))

	// Jika ada kesalahan dalam konversi, langsung return false
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return false
	}

	firstSum := firstDigit1 + firstDigit2
	lastSum := lastDigit1 + lastDigit2

	if firstSum != lastSum {
		return false
	}

	// Cek digit tengah (semua harus lebih besar dari 5)
	var middleDigits string
	if length == 6 {
		middleDigits = ticket[2:4]
	} else {
		middleDigits = ticket[2:6]
	}

	for _, digit := range middleDigits {
		digitValue, err := strconv.Atoi(string(digit))
		if err != nil || digitValue <= 5 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah mahasiswa (N): ")
	fmt.Scan(&n)

	validCount := 0
	invalidCount := 0

	// Memproses setiap tiket
	for i := 0; i < n; i++ {
		var ticket string
		fmt.Print("Masukkan nomor tiket: ")
		fmt.Scan(&ticket)

		if isValidTicket(ticket) {
			validCount++
		} else {
			invalidCount++
		}
	}

	// Menampilkan hasil
	fmt.Printf("Tiket valid: %d\n", validCount)
	fmt.Printf("Tiket tidak valid: %d\n", invalidCount)
}