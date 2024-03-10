package days

import (
	"fmt"
	"project/funcs"
)

func Day13() {
	fmt.Println("1.QR-BANK")
	var x int 
	fmt.Scan(&x) 
	switch x {
	case 1:
		funcs.Qr_main()
	default:
		fmt.Println("Please enter the valid number!" )
	}
}