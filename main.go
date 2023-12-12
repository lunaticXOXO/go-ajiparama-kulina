package main

import (
	"github.com/go-ajiparama-kulina/controller"
)

func main(){

	input_1 := []int{10,20,20,10,10,30,50,10,20}
	input_2a := 8
	input_2b := "UDDDUDUU"
	input_3 := "1345679"
	input_4 := 100
	
	controller.CountPair(input_1)
	controller.CountValleys(input_2a,input_2b)
	controller.CreateUnique(input_3)
	controller.CountLamps(input_4)


}