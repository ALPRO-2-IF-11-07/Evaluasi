package main

import "fmt"

func main() {
	var n int
	var m int

	fmt.Print("Masukan Pembilang (n): ")
	fmt.Scanln(&n)
	fmt.Print("Masukan Pembilang (n): ")
	fmt.Scanln(&m)

	hasilPembagian := n / m

	hasilModulus := n % m

	fmt.Println("Hasil pembagian dari ", n, " dengan ", m, " adalah: ", hasilPembagian)
	fmt.Println("Sisa dari pembagian ", n, " dengan ", m, " adalah: ", hasilModulus)

}
