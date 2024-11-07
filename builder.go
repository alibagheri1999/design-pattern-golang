package main

import "log"

type car struct {
	brand string
	model string
	color string
	year int
}

func (c *car) GetBrand() string {
	return c.brand
}

func (c *car) GetModel() string {
	return c.model
}

func (c *car) GetColor() string {
	return c.color
}

func (c *car) GetYear() int {
	return c.year
}

type CarBuilder struct {
	car car
}

func NewCarBuilder() *CarBuilder  {
	return &CarBuilder{}
}

func (b *CarBuilder)SetBrand(brand string) *CarBuilder {
	b.car.brand = brand
	return b
}

func (b *CarBuilder)SetModel(model string) *CarBuilder {
	b.car.model = model
	return b
}

func (b *CarBuilder)SetColor(color string) *CarBuilder {
	b.car.color = color
	return b
}

func (b *CarBuilder)SetYear(year int) *CarBuilder {
	b.car.year = year
	return b
}

func (b *CarBuilder) Build() car {
	return b.car
} 

func (b *CarBuilder) SampleCar() car  {
	return b.SetBrand("Sample Car").SetModel("New One").SetColor("Gray").SetYear(1999).Build()
}

func builder()  {
	sample := NewCarBuilder().SampleCar()
	log.Println(sample.GetBrand(), sample.GetModel(), sample.color, sample.GetYear())
}
