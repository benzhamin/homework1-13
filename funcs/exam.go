package funcs

import (
	"fmt"
	"math/rand"
)

func Type_exam(student *Student) {
	x := rand.Intn(2) + 1

	switch x {
	case 1:
		student.Exam = "Math"
	case 2:
		student.Exam = "English"
	case 3:
		student.Exam = "Russian"
	}

}

func Start_exam(student *Student) {
	switch student.Exam {
	case "Math":
		student.Otsenkaexamina = rand.Intn(4) + 1
		if student.Otsenki.Math > student.Otsenkaexamina {
			student.Otsenkaexamina++
			fmt.Println("Teacher increase your exam grade because of your leasson grades !!")
		}

		if student.Xuliganstvo.Absents > 5 || student.Xuliganstvo.Fights > 5 {
			student.Otsenkaexamina--
			fmt.Println("Teacher Decrase you exam grade beacuse of your Xuliganstvo")
		}

	case "English":
		student.Otsenkaexamina = rand.Intn(4) + 1
		if student.Otsenkaexamina < 5 {
			student.Otsenkaexamina = 5
			fmt.Println("English Teacher incrase your grade to 5 because she like you :) ")
		}
	case "Russian":
		student.Otsenkaexamina = rand.Intn(4) + 1
		student.Otsenkaexamina -- ;
		fmt.Println("Russian Teacher dicrase you grade beacuse you talk in her leassons :( ")
	}
}
