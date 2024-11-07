package main

import (
	"fmt"
	"log"
)

type Coffee interface {
	Name() string
	Price() int
}

type SimpleCoffee struct{}

func (s SimpleCoffee) Name() string {
	return "Coffee"
}

func (s SimpleCoffee) Price() int {
	return 50000
}

type WithMilk struct{
	coffee Coffee
}

func (w WithMilk) Name() string {
	return fmt.Sprintf("%s Milk", w.coffee.Name())
}

func (w WithMilk) Price() int {
	return w.coffee.Price() + 10000
}

type WithSuger struct{
	coffee Coffee
}

func (w WithSuger) Name() string {
	return fmt.Sprintf("%s Suger", w.coffee.Name())
}

func (w WithSuger) Price() int {
	return w.coffee.Price() + 5000
}

func decorator() {
	coffee := SimpleCoffee{}
	coffeeWithSuger := WithSuger{coffee}
	latte := WithMilk{coffee}
	latteWithSuger := WithSuger{latte}

	items := []Coffee{
		coffee, coffeeWithSuger, latte, latteWithSuger,
	}

	for _, coffee := range items {
		log.Printf("%-*s and Price %d T\n",30 , coffee.Name(), coffee.Price())
	}
}