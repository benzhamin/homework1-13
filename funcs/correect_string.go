package funcs

import (
	"fmt"
	"unicode/utf8"
)

func c_string( s string, ln *int ) (ans string) {
	for i := 0; i < *ln; i++ {
		if s[i] == '+' {
			ans += " " ;
		}else {
			ans += string(s[i])
		}
		
	}
	return ans 
}

func Corect_string() {
	var s string = "My+name+is+Benjamin"
	ln := utf8.RuneCountInString(s)
	fmt.Println(c_string(s,&ln))
}
