package controller


import (
	"fmt"

)

func CountPair(array []int){
	fmt.Println("Answered number 1 ")
	pair := make(map[int]struct{})
	countPair := 0

	for _, num := range array{
		if _, exists := pair[num];exists{
			countPair++
			delete(pair,num)
		}else{
			pair[num] = struct{}{}
		}
	}

	fmt.Println("total pair : ", countPair)
	
}

func CountValleys(n int, steps string){
	fmt.Println("Answered number 2 ")
	altitude := 0
	valley := 0

	for i := 0; i < n; i++{
		if steps[i] == 'U'{
			altitude++
		}else{
			altitude--
		}
		if altitude == 0 && steps[i] == 'U' {
			valley++
		}
	}
	fmt.Println("valleys : ",valley)
	
}

func CreateUnique(inputString string) {
	fmt.Println("Answered Number 3")
	zero := "0"
	length := len(inputString)

	for i := 0; i < length; i++ {
		for j := length; j > i; j-- {
			if j == length {
				fmt.Print(string(inputString[i]))
			} else {
				fmt.Print(zero)
			}
		}
		fmt.Print("\n")
	}
	
}

func CountLamps(trips int) {
	fmt.Println("Answered Number 4")
	lamps := make([]int, 101)
	count := 0
	for i := 1; i <= trips; i++ {
		for j := i; j <= trips; j += i {
			lamps[j] = 1 - lamps[j]
		}
	}

	for k := 1; k <= trips; k++ {
		count += lamps[k]
	}
	fmt.Println("lamps : ",count)
	
}

