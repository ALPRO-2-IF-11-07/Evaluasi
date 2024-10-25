package main

import (
	"fmt"
	"strconv"
)

func isValidTicket(ticket string) bool {

	n := len(ticket)
	if n != 6 && n != 8 {
		return false
	}

	firstTwoSum, _ := strconv.Atoi(ticket[0:2])
	lastTwoSum, _ := strconv.Atoi(ticket[n-2 : n])
	if firstTwoSum != lastTwoSum {
		return false
	}

	for i := 2; i < n-2; i++ {
		digitTengah, _ := strconv.Atoi(string(ticket[i]))
		if digitTengah <= 5 {
			return false
		}
	}

	return true
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah mahasiswa (N): ")
	fmt.Scanln(&N)

	hitungValid := 0
	hitungInvalid := 0

	for i := 0; i < N; i++ {
		var ticket string
		fmt.Printf("Masukkan nomor tiket: ")
		fmt.Scanln(&ticket)

		if isValidTicket(ticket) {
			hitungValid++
			fmt.Println("Tiket valid")
		} else {
			hitungInvalid++
			fmt.Println("Tiket tidak valid")
		}
	}


	fmt.Println("Jumlah tiket valid:", hitungValid)
	fmt.Println("Jumlah tiket tidak valid:", hitungInvalid)
}
