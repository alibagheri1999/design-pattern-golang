package main

import (
	"errors"
	"log"
	"math/rand"
)

type IPaymentGateway interface {
	Pay(amount int) error
}

type SamanPayment struct{}

func (s *SamanPayment) StartTrnsaction(amount int, uniqueID float64) {
	log.Println("Start Saman Payment")
}

type PaymentConfig struct {
	ClientID int
	CleintSecret string
}

type AsanPardakhtPayment struct {
	config PaymentConfig
}

func (a *AsanPardakhtPayment) Process(ClientID int, clientSecret string, amount float64) {
	log.Println("Start Asan Pardakht Payment")
}

type SamanPaymenentAdapter struct {
	Payment SamanPayment
}

func generateRandomID() float64 {
 return rand.Float64()
}

func (a *SamanPaymenentAdapter) Pay(amount int) error {
	uniqueID := generateRandomID()
	a.Payment.StartTrnsaction(amount, uniqueID)
	log.Println("Saman payment completed")
	return nil
}

type AsanPardakhtPaymentAdapter struct {
	Payment AsanPardakhtPayment
}

func (a *AsanPardakhtPaymentAdapter) Pay(amount int) error {
	a.Payment.Process(a.Payment.config.ClientID, a.Payment.config.CleintSecret, float64(amount))
	log.Println("Asan pardakht payment completed")
	return nil
}

var config = PaymentConfig{
	ClientID: 123456,
	CleintSecret: "CleintSecret",
}

func PaymentAdapter(paymentType string) (IPaymentGateway, error) {
	switch paymentType {
	case "saman":
		samanPayment := SamanPayment{}
		adapter := SamanPaymenentAdapter{Payment: samanPayment}
		return &adapter, nil
	case "asanpardakht":
		asanPardakhtPayment := AsanPardakhtPayment{config}
		adapter := AsanPardakhtPaymentAdapter{Payment: asanPardakhtPayment}
		return &adapter, nil
	default:
		return nil, errors.New("Invalid Payment type")
	}
}

func adapter()  {
	paymentGateway, err := PaymentAdapter("saman")
	if err != nil {
		log.Println(err)
		return
	}
	err = paymentGateway.Pay(100)
	if err != nil {
		log.Println(err)
		return
	}
}