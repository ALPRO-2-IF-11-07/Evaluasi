package main

import (
        "fmt"
        "math"
)
//HANAH NUR AZIZAH 2311102312
func main() {
        var num int
        fmt.Print("Masukkan bilangan bulat positif lebih dari 100: ")
        fmt.Scan(&num)

        digitCount := int(math.Log10(float64(num))) + 1

        partLength := digitCount / 3
        if digitCount%3 != 0 {
                partLength++
        }

        part1 := num / int(math.Pow10(digitCount-partLength))
        num %= int(math.Pow10(digitCount-partLength))
        part2 := num / int(math.Pow10(digitCount-2*partLength))
        part3 := num % int(math.Pow10(digitCount-2*partLength))

        average := (part1 + part2 + part3) / 3

        fmt.Println("Hasil pemotongan:")
        fmt.Println(part1, part2, part3)
        fmt.Println("Rata-rata dari ketiga bilangan:")
        fmt.Println(average)
}