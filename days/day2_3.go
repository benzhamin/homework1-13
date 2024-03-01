package days

import (
	"fmt"
	"project/funcs"
)

func Day2_3() {
	var x int 
	fmt.Println("Chose: \n1: Anniversary \n2: Area \n3: Perimeter")
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
