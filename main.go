package main

import (
	"github.com/go-ajiparama-kulina/controller"
	"fmt"
)
func main() {
	var input_2a int
	var input_2b string
	var input_3 string
	var input_4 int
	var input_soal string
	var size int

	fmt.Println("Input nomor soal : (1-4 ?)")
	fmt.Scanln(&input_soal)

	switch input_soal {
	case "1":
		fmt.Println("Input jumlah ukuran array : ")
		fmt.Scanln(&size)
		input_1 := make([]int, size)

		for i := 0; i < size; i++ {
			fmt.Printf("Enter element %d: ", i+1)
			fmt.Scanln(&input_1[i])
		}
		fmt.Println("Array input_1:", input_1)
		controller.CountPair(input_1)

	case "2":
		fmt.Println("Input jumlah langkah : ")
		fmt.Scanln(&input_2a)
		fmt.Println("Input jalur (U/D)")
		fmt.Scanln(&input_2b)
		controller.CountValleys(input_2a, input_2b)

	case "3":
		fmt.Println("Input angka : ")
		fmt.Scanln(&input_3)
		controller.CreateUnique(input_3)

	case "4":
		fmt.Println("Input jumlah (1-100) : ")
		fmt.Scanln(&input_4)
		controller.CountLamps(input_4)

	default:
		fmt.Println("Input invalid")
	}
}