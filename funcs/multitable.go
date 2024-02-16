package funcs

import (
	"fmt"
)

func MultiTable() {
	var x, y int
	fmt.Println("Multi table from x to y. Please enter X & Y ")
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		fmt.Println("Table", i)
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}
}
