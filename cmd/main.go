package main

import (
	"fmt"
	"log"
	"project/funcs"
)

func do(day int) {
	fmt.Println("Chose: \n1: Anniversary \n2: Area \n3: Perimeter")
	var x int
	fmt.Scan(&x)
	switch x {
	case 1:
		funcs.Anniversary()
	case 2:
		funcs.Area()
	case 3:
		funcs.Perimeter()
	default:
		fmt.Println("Please enter valid number")
	}
}

func main() {

	log.Println("Program Started")

	for {
		fmt.Println("Choose day: from 1 to 4 ")
		var day int
		fmt.Scan(&day)

		switch day {
		case 1:
			fmt.Println("Znakomvstvo...")
		case 2:
			do(day)
		case 3:
			do(day)
		case 4:
			fmt.Println("1: Multi Table \n2: Fibonachi")
			var x int
			fmt.Scan(&x)
			switch x {
			case 1:
				funcs.MultiTable()
			case 2:
				funcs.Fibonacci()
			default:
				fmt.Println("Please enter the valid number")
			}

		default:
			fmt.Println("Please enter the valid number")
		}

	}

} 
