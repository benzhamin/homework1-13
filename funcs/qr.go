package funcs

import "fmt"

type Alif struct {
	Name  string
	Money float64
}

type Dc struct {
	Name  string
	Money float64
}

type Humo struct {
	Name  string
	Money float64
}

type QR interface {
	Check_balance()
	Add_balance(float64)
}

func qr_chek(user QR) {
	user.Check_balance()
}
func qr_add(user QR) {
	user.Add_balance(100)
}

// from user1 to user2
func qr_transfer(user1 QR, user2 QR, money float64) {
	user1.Add_balance(-money)
	user2.Add_balance(money)
	fmt.Println("Transfer is Done!!!")
}

func Qr_main() {
	var (
		user_humo Humo
		user_alif Alif
		user_dc   Dc
	)
	user_humo.Name = "Humo_Benjamin"
	user_alif.Name = "Alif_Benjamin"
	user_dc.Name = "Dc_Benjamin"

	qr_add(&user_humo)
	qr_add(&user_alif)
	qr_add(&user_dc)
	qr_chek(&user_humo)
	qr_chek(&user_alif)
	qr_chek(&user_dc)
	qr_transfer(&user_alif, &user_humo, 50)
	qr_chek(&user_humo)
	qr_chek(&user_alif)
	qr_chek(&user_dc)
}

func (user *Alif) Check_balance() {
	fmt.Println(user.Name, user.Money)
}

func (user *Alif) Add_balance(add float64) {
	user.Money += add
}

func (user *Dc) Check_balance() {
	fmt.Println(user.Name, user.Money)
}

func (user *Dc) Add_balance(add float64) {
	user.Money += add
}

func (user *Humo) Check_balance() {
	fmt.Println(user.Name, user.Money)
}

func (user *Humo) Add_balance(add float64) {
	user.Money += add
}
