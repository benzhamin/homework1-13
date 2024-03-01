package days

import (
	"fmt"
	"math/rand"
	"project/funcs"
)

func Day5() {
	var Student funcs.Student = funcs.Student{Name: "Satoru Gojo", Class: rand.Intn(10) + 1,
		Otsenki: funcs.Otsenki{Math: rand.Intn(4)+1, English: rand.Intn(4)+1, Russian: rand.Intn(4)+1}, 
		Xuliganstvo: funcs.Xuliganstvo{Absents: rand.Intn(10), Fights: rand.Intn(1000)}}

	funcs.Type_exam(&Student) 
	fmt.Println("Todays exam is", Student.Exam)
	fmt.Println("---------------------" )
	funcs.Start_exam(&Student)

	fmt.Println(Student.Name, "Your exam grade is",Student.Otsenkaexamina) 
	
}
 