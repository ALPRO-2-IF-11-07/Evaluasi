package main

import "fmt"

func potongBilangan(bilangan int) (int, int) {

	bagian1 := bilangan / 100
	bagian2 := bilangan % 100

	return bagian1, bagian2
}

func main() {
	var bilangan int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	bagian1, bagian2 := potongBilangan(bilangan)

	if bilangan <= 100 {
		fmt.Println("Bilangan harus lebih besar dari 100")
	} else {
		if bilangan%2 == 0 {
			fmt.Println(bagian1)
			fmt.Println(bagian2)
		}

	}
}
