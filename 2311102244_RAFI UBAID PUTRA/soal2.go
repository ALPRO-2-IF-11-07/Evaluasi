package main

import "fmt"

func main() {
	var peserta int
	var tiket int

	fmt.Print("Jumlah Peserta : ")
	fmt.Scanln(&peserta)

	if peserta < 1 {
		fmt.Println("Jumlah peserta minimal 1 ")
	} else {

		for i := 0; i < peserta; i++ {
			fmt.Print("Masukan nomor tiket peserta acara ", i+1, ": ")
			fmt.Scan(&tiket)
			if tiket%2 == 0 {
				fmt.Println("Hadiah Utama")
			} else if tiket%2 != 0 {
				fmt.Println("Hadiah Sembako")
			} else {
				fmt.Println("Hadiah Konsol")
			}
		}

	}

}
