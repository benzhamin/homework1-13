package funcs

import "fmt"

func fib(pre, cur, pos int, x, y *int) {
	if pos >= *x && pos <= *y {
		fmt.Printf("index %d fib %d \n", pos, cur)
	}
	if pos > *y  {
		return
	}
	fib(cur, cur+pre, pos+1, x, y)
}

func Fibonacci() {
	fmt.Println("Fibonacci index from x yo y")
	var x, y int
	fmt.Scan(&x, &y)
	X := &x 
	Y := &y
	fib(0, 1, 1, X, Y)
}
