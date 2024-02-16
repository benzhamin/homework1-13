package funcs

import "fmt"

func Perimeter() {
	var x, y int
	fmt.Println("Please enter the sides: x & y")
	fmt.Scan(&x, &y)

	fmt.Println("Perimeter is: ", 2*(x+y))

}
