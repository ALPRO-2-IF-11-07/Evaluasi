// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var num int
// 	fmt.Print("Masukkan bilangan bulat positif lebih dari 100: ")
// 	fmt.Scan(&num)

// 	digit1 := num / 100
// 	digit2 := (num / 10) % 10
// 	digit3 := num % 10

// 	fmt.Println("Hasil pemotongan:")
// 	fmt.Println(digit1, digit2, digit3)

// 	rataRata := (digit1 + digit2 + digit3) / 3
// 	fmt.Println("Rata-rata dari ketiga bilangan:", rataRata)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func cekHadiah(tiket int) string {
// 	strTiket := strconv.Itoa(tiket)
// 	genap := true
// 	ganjil := true

// 	for _, digit := range strTiket {
// 		d := int(digit - '0')
// 		if d%2 == 0 {
// 			ganjil = false
// 		} else {
// 			genap = false
// 		}
// 	}

// 	if genap {
// 		return "Hadiah Utama"
// 	} else if ganjil {
// 		return "Hadiah Sembako"
// 	} else {
// 		return "Hadiah Konsol"
// 	}
// }

package main

import (
	"fmt"
	"strconv"
)

func cekHadiah(tiket int) string {
	strTiket := strconv.Itoa(tiket)
	genap := true
	ganjil := true

	for _, digit := range strTiket {
		d := int(digit - '0')
		if d%2 == 0 {
			ganjil = false
		} else {
			genap = false
		}
	}

	if genap {
		return "Hadiah Utama"
	} else if ganjil {
		return "Hadiah Sembako"
	} else {
		return "Hadiah Konsol"
	}
}

func main() {
	var N int
	var hadiahUtama, hadiahSembako, hadiahKonsol int

	fmt.Print("Masukkan jumlah peserta acara: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		var tiket int
		fmt.Printf("Masukkan nomor tiket peserta %d: ", i)
		fmt.Scan(&tiket)

		hadiah := cekHadiah(tiket)
		fmt.Println(hadiah)

		switch hadiah {
		case "Hadiah Utama":
			hadiahUtama++
		case "Hadiah Sembako":
			hadiahSembako++
		case "Hadiah Konsol":
			hadiahKonsol++
		}
	}

	fmt.Println("Jumlah:")
	fmt.Println("Hadiah Utama:", hadiahUtama)
	fmt.Println("Hadiah Sembako:", hadiahSembako)
	fmt.Println("Hadiah Konsol:", hadiahKonsol)
}

// package main

// import (
// 	"fmt"
// )

// func pembagian(n int, m int) int {
// 	if n < m {
// 		return 0
// 	}
// 	return 1 + pembagian(n-m, m)
// }

// func sisaPembagian(n int, m int) int {
// 	if n < m {
// 		return n
// 	}
// 	return sisaPembagian(n-m, m)
// }

// func main() {
// 	var n, m int

// 	fmt.Print("Masukkan bilangan pembilang (n): ")
// 	fmt.Scan(&n)
// 	fmt.Print("Masukkan bilangan penyebut (m): ")
// 	fmt.Scan(&m)

// 	hasil := pembagian(n, m)
// 	sisa := sisaPembagian(n, m)

// 	fmt.Printf("Hasil dari pembagian %d dengan %d adalah: %d\n", n, m, hasil)
// 	fmt.Printf("Sisa dari pembagian adalah: %d\n", sisa)
// }
