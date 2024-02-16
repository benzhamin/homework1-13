package funcs

import (
	"fmt"
	"log"
	"math"
)

func Anniversary() {
	log.Println("Program started")
	var age int
	fmt.Println("Please enter your age : ")
	fmt.Scan(&age)
	fmt.Println("до следуйщиво юбилея осталось: ", math.Ceil(float64(age)/10)*10-float64(age), " год")
	log.Println("Program ended")
}
