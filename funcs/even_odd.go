package funcs

import "fmt"

func Even_Odd() {
	A := []int{1, 2, 3, 4, 5, 6}
	odd := []int{}
	even := []int{}

	for _, x := range A {
		if x % 2 == 0{
			even = append(even, x)
		}else {
			odd = append(odd, x)
		}
	}
	fmt.Print("Evens: ")
	for _, v := range even {
		fmt.Print(v, " ")
	}
	fmt.Print("\nOdds: ")
	for _, v := range odd {
		fmt.Print(v," ")
	}
	fmt.Println()
}
