package days

import (
	"fmt"
	"project/funcs"
)

func Day4() {
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
}
