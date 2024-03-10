package days

import (
	"fmt"
	"project/funcs"
)

func Day11() {
	fmt.Print("1.CONV\n2.Sdvig3\n3.Correct_String\n")
	var x int
	fmt.Scan(&x)
	switch x {
	case 1:
		funcs.Conv()
	case 2:
		funcs.Sdvig3()
	case 3:
		funcs.Corect_string()
	default:
		fmt.Println("Please enther the valid numer")
	}
}
