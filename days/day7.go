package days

import (
	"fmt"
	"project/funcs"
)

func Day7() {
	fmt.Println("1. Reverse Massive\n2. Even&odd")
	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		funcs.Reverse()
	case 2:
		funcs.Even_Odd()
	default:
		fmt.Println("Please enter the valid number")
	}
}
