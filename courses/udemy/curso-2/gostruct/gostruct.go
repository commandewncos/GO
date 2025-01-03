package gostruct

import "time"

type Address struct {
	City, Street string
	Number       int
}

type Person struct {
	Name       string
	Nascimento time.Time
	Email      string
	Address
}

func (p Person) Age() int {
	return time.Now().Year() - p.Nascimento.Year()
}

func (p *Person) ChangeEmail(newEmail string) {
	p.Email = newEmail

}
