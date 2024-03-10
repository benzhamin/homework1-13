package funcs

import (
	"fmt"
	"unicode/utf8"
)

func shifravanie(s string, ln *int) (ans string ){
	for i := 0; i < *ln; i++ {
		x := s[i] + 3 
		ans += string(x)
	}
	return ans 
}

func razshif(s string, ln *int ) ( ans string ) {
	for i := 0; i < *ln; i++ {
		x := s[i] - 3 
		ans += string(x) 
	}
	return ans 
}

func Sdvig3() {

	var s string = "My name is Benjamin"
	ln := utf8.RuneCountInString(s)
	ss := shifravanie(s,&ln)
	fmt.Println(ss)
	fmt.Println(razshif(ss,&ln))

}
