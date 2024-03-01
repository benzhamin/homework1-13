package funcs

import (
	"fmt"
	"math"
)

func Anniversary() {
	var age int
	fmt.Println("Please enter your age : ")
	fmt.Scan(&age)
	fmt.Println("до следуйщиво юбилея осталось: ", math.Ceil(float64(age)/10)*10-float64(age), " год")
}
