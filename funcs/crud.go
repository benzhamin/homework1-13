package funcs

import (
	"fmt"
)

type worker struct {
	Name    string
	Age     int
	Exp     float64
	Married bool
}

func create(name string, age int, exp float64, married bool) (user worker) {
	user.Name = name
	user.Age = age
	user.Exp = exp
	user.Married = married
	return user
}
func update(user worker) worker {
	fmt.Printf("Name: %v Age: %v Exp: %v Married: %v\n", user.Name, user.Age, user.Exp, user.Married) // not working
	fmt.Println("Select what do you want to update!!!")
	var update_user string
	fmt.Scan(&update_user)
	fmt.Println("To waht do you want to change")
	
	switch update_user {
	case "Name":
		var x string 
		fmt.Scan(&x) 
		user.Name = x
	case "Age":
		var x int 
		fmt.Scan(&x)
		user.Age = x 
	case "Exp":
		var x float64
		fmt.Scan(&x)
		user.Exp = x
	case "married":
		var x bool 
		fmt.Scan(&x) 
		user.Married = x
	default:
		fmt.Println("Wrong input")
	}

	return user
}
func delate_map( key_map string, mp *map[string]worker ) {
	delete(*mp, key_map)
}
func read( user worker ) {
	fmt.Printf("Name: %v Age: %v Exp: %v Married: %v\n", user.Name, user.Age, user.Exp, user.Married) // not working
}

func Crud() {
	mp := make(map[string]worker)
	mp["user1"] = create("Benjamin Franklin", 19, 3.5, false)
	fmt.Println(mp["user1"])
	mp["user1"] = update(mp["user1"])
	read(mp["user1"])
	delate_map("user1", &mp)
	fmt.Println(mp)
}
