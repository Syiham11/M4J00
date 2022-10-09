package main

import "fmt"

func DeretAritmatika() []int {
	var (
		num1, num2, n int
	)
	fmt.Println("Masukan angka deret pertama : ")
	fmt.Scanln(&num1)
	fmt.Println("Masukan angka deret kedua : ")
	fmt.Scanln(&num2)
	fmt.Println("Berapa angka deret yang ingin ditampilkan ? ")
	fmt.Scanln(&n)
	fmt.Println(num1, num2, n)

	deretAritmatik := []int{num1, num2}
	difference := num2 - num1
	current := num2

	for i := 0; i < (n - 2); i++ {
		current += difference
		deretAritmatik = append(deretAritmatik, current)
	}
	return deretAritmatik
}

func main() {
	fmt.Println(DeretAritmatika())
}
