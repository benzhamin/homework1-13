package funcs

import "fmt"

func Area() {
	var x, y int
	fmt.Println("Please enter the sides: x & y")
	fmt.Scan(&x, &y)

	fmt.Println("Area is: ", x*y)

}
