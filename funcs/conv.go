package funcs

import (
	"fmt"
	"unicode/utf8"
)

func convtobyte(s string, ln *int ) ( mas []int ) {
	for i := 0; i < *ln; i++ {
		mas = append(mas, int(s[i])) 
	}
	return mas 
}

func convtostring(mas []int, ln *int) ( s string ) {
	for i := 0; i < *ln; i++ {
		s += fmt.Sprint(mas[i]) 
	}
	return s 
}

func Conv() {
	s := "Hello World..." 
	ln := utf8.RuneCountInString(s)

	mas := convtobyte(s, &ln) 
	fmt.Println(mas) 
	fmt.Println(convtostring(mas,&ln))
}
