package funcs

import "fmt"

func Reverse() {
	fmt.Println("Please enter the size of Massive " ) 
	var n int 
	var A[100] string
	fmt.Scan(&n) 

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
		if i > n/2 {
			A[i], A[n-i-1] = A[n-i-1], A[i]
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
/// 1 2 3 4 5