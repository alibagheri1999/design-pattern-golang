package main

import "log"

type PaymentType = uint8

const (
	TypeSaman PaymentType = iota + 1
	TypeMellat
)

type PaymentGateway interface {
	Pay(amount int)
}

type saman struct{}

func (s saman) Pay(amount int)  {
	// do some payment
	log.Println("Saman Pay")
}

func NewSaman() PaymentGateway {
	return saman{}
}

type mellat struct{}

func (m mellat) Pay(amount int) {
	// do some payment
	log.Println("Mellat Pay")
}

func NewMellat() PaymentGateway {
	return mellat{}
}

func PaymentFactory(t PaymentType) PaymentGateway {
	switch t {
	case TypeMellat:
		return NewMellat()
	default: 
		return NewSaman()
	}
}

func CheckOut(t PaymentType, amount int)  {
	payment := PaymentFactory(t)
	payment.Pay(amount)
}

func factory()  {
	// it can be TypeSaman or Type Mellat
	CheckOut(TypeSaman, 1000)
}