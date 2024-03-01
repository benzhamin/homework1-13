package main

import (
	"fmt"
	"log"
	"project/days"
)

func main() {

	log.Println("Program Started")

	fmt.Println("Choose day: from 1 to 7 ")
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
	default:
		fmt.Println("Please enter the valid number")
	}
	log.Println("Program ended")

}
