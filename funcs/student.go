package funcs

type Otsenki struct {
	Math    int
	English int
	Russian int
}

type Xuliganstvo struct {
	Absents int // how many proguli
	Fights int
}

type Student struct {
	Name    string
	Class   int /// from 1 to 11
	Otsenki Otsenki
	Xuliganstvo Xuliganstvo
	Exam string 
	Otsenkaexamina int 
}
