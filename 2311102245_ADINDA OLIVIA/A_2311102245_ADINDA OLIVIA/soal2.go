package main

import (
	"fmt"
)

// fungsi untuk memeriksa apakah semua digit dalam nomor tiket adalah ganjil
func semuaDigitganjil(nomor int) bool {
	for nomor > 0 {
		digit := nomor % 10
		if digit%2 == 0 { // perbaikan di sini, periksa ganjil
			return false
		}
		nomor /= 10
	}
	return true
}

// fungsi untuk memeriksa apakah semua digit dalam nomor tiket sama
func semuaDigitsama(nomor int) bool {
	digitSama := nomor % 10
	nomor /= 10
	for nomor > 0 {
		digit := nomor % 10
		if digit != digitSama {
			return false
		}
		nomor /= 10
	}
	return true
}

// fungsi untuk memeriksa apakah semua digit dalam nomor tiket adalah genap
func semuaDigitgenap(nomor int) bool {
	for nomor > 0 {
		digit := nomor % 10
		if digit%2 != 0 {
			return false
		}
		nomor /= 10
	}
	return true
}

// fungsi untuk menentukan jenis hadiah berdasarkan nomor tiket
func tentukanHadiah(nomor int) string {
	// periksa apakah semua digit sama dan genap
	if semuaDigitsama(nomor) && semuaDigitgenap(nomor) {
		return "hadiah utama"
	}
	// periksa apakah semua digit adalah ganjil
	if semuaDigitganjil(nomor) {
		return "hadiah sembako"
	}

	// jika tidak memenuhi kedua syarat di atas, beri hadiah konsol
	return "hadiah konsol"
}

func main() {
	// input dari pengguna
	var N int
	fmt.Print("Masukkan jumlah anggota: ")
	fmt.Scan(&N)

	// inisialisasi penghitung untuk setiap jenis hadiah
	jumlahHadiahUtama := 0
	jumlahHadiahSembako := 0
	jumlahHadiahKonsol := 0

	// Proses setiap nomor tiket
	for i := 0; i < N; i++ {
		var nomor int
		fmt.Printf("Masukkan nomor tiket anggota ke-%d: ", i+1)
		fmt.Scan(&nomor)

		// proses setiap nomor tiket
		jenisHadiah := tentukanHadiah(nomor)
		fmt.Println(jenisHadiah)

		// update penghitung berdasarkan jenis hadiah
		switch jenisHadiah {
		case "hadiah utama":
			jumlahHadiahUtama++
		case "hadiah sembako":
			jumlahHadiahSembako++
		case "hadiah konsol":
			jumlahHadiahKonsol++
		}
	}

	// tampilkan total jumlah setiap jenis hadiah
	fmt.Printf("jumlah: hadiah utama = %d, hadiah sembako = %d, hadiah konsol = %d\n", jumlahHadiahUtama, jumlahHadiahSembako, jumlahHadiahKonsol)
}
