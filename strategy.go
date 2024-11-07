package main

import "log"

type shoppingMethodType uint8

const (
	shippingMethodPeyk shoppingMethodType = iota + 1
	shoppingMethodPost 
)

type ShppingStrategy interface {
	Ship()
}

type PeykShpping struct{}

func (p *PeykShpping) Ship()  {
	log.Println("PeykShpping")
}

type PostkShpping struct{}

func (p *PostkShpping) Ship()  {
	log.Println("PostkShpping")
}

type Order struct{
	shppingStrategy ShppingStrategy
}

func (o *Order) SetStrategy(strategy ShppingStrategy) {
	o.shppingStrategy = strategy
}

func (o *Order) Ship() {
	o.shppingStrategy.Ship()
}

func strategy()  {
	order := &Order{}

	shoppingMethod := shippingMethodPeyk

	switch shoppingMethod {
	case shippingMethodPeyk:
		order.SetStrategy(&PeykShpping{})
	case shoppingMethodPost:
		order.SetStrategy(&PostkShpping	{})
	default:
		log.Println("Invalid Shipping Method")
	}
	order.Ship()
}