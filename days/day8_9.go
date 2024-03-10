package days

import (
	"fmt"
	"project/funcs"
)

func Day8_9() {
	fmt.Print("1.CRUD\n")
	var x int 
	fmt.Scan(&x) 
	switch x {
	case 1 :
		funcs.Crud()
	default:
		fmt.Println("Please enter the valid number")
	}
}