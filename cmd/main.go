package main

import (
	"fmt"
	"log"
	"project/days"
)

func main() {

	log.Println("Program Started")

	fmt.Println("Choose day: from 1 to 13 ")
	var day int
	fmt.Scan(&day)

	switch day {
	case 1:
		days.Day1()
	case 2:
		days.Day2_3()
	case 3:
		days.Day2_3()
	case 4:
		days.Day4()
	case 5: 
		days.Day5()
	case 6:
		days.Day6()
	case 7:
		days.Day7()
	case 8,9:
		days.Day8_9()
	case 10:
		days.Day10()
	case 11:
		days.Day11() 
	case 12:
		days.Day12()
	case 13:
		days.Day13()
	default:
		fmt.Println("Please enter the valid number")
	}
	log.Println("Program ended")

}
