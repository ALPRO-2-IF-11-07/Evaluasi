package main

import "fmt"


func sumEvenRecursive(nums []int, indeks int, sum int) int {
	if indeks == len(nums) {
		return sum
	}

	if nums[indeks] > 0 && nums[indeks]%2 == 0 {
		sum += nums[indeks]
	}

	return sumEvenRecursive(nums, indeks+1, sum)
}

func main() {
	var num int
	nomor := []int{}

	fmt.Println("Masukkan bilangan bulat (bilangan negatif sebagai sentinel):")
	for {
		fmt.Scanln(&num)
		if num < 0 {
			break
		}
		nomor = append(nomor, num)
	}

	sumRecursive := sumEvenRecursive(nomor, 0, 0)
	fmt.Printf("Jumlah angka genap (metode rekursif): %d\n", sumRecursive)
}
