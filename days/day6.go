package days

import (
	"fmt"
	"log"
	"math/rand"
	"project/funcs/customer"
	"project/funcs/operations"
)

func Day6() {
	//Тестовый проект - "Дурная больница"
	//Авторы - студенты HumoLab
	var Customers []customer.Patient = []customer.Patient{{Name: "Гордон Фриманович Ломов", Age: rand.Intn(126) + 5, Order: 0, Organs: customer.Organs{Leg: 2, Hands: 2, Livers: 2, Heart: false}},
		{Name: "Benjamin", Age: rand.Intn(123) + 5, Order: 0, Organs: customer.Organs{Leg: 2, Hands: 2, Livers: 2, Heart: false} }}

	for _, Customer := range Customers {
		log.Println("Я зашёл забрать чек")
		operations.TakeOrder(&Customer)
		log.Println("Я забрал чек. \nМой номер:", Customer.Order)
		log.Println("Я пошёл на операцию.")
		operations.BeginOperation(&Customer)
		//Проверка результата после операции
		if !Customer.Organs.Heart {
			log.Printf("Я Вернулся! Но что-то со мной не так: %+v", Customer)
		} else {
			log.Printf("Я Вернулся! И чувствую себя прекрасно!: %+v", Customer)
		}
		fmt.Println()
	}
}
