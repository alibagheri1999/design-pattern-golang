package main

import "log"

type User struct {
	Name string
	Email string
	Phone string
}


type Observer interface {
	Update(user User)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) UnSubscribe(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]... )
			break
		}
	}
}

func (s *Subject) Notify(user User) {
	for _, observer := range s.observers {
		observer.Update(user)
	}
}

type SMSService struct{}

func (s *SMSService) Update(user User) {
	log.Println("Sending SMS", user)
}

type EmailService struct{}

func (s *EmailService) Update(user User) {
	log.Println("Sending Email", user)
}

func observer() {
	userRegistration := &Subject{}

	smsService := &SMSService{}
	emailService := &EmailService{}

	userRegistration.Subscribe(smsService)

	userRegistration.Notify(User{
		Name: "Al1",
		Email: "ali1.bagheri.software@gmail.com",
		Phone: "123456789",
	})

	userRegistration.Subscribe(emailService)

	userRegistration.Notify(User{
		Name: "Ali2",
		Email: "ali2.bagheri.software@gmail.com",
		Phone: "987654321",
	})

	userRegistration.UnSubscribe(smsService)

	userRegistration.Notify(User{
		Name: "Ali3",
		Email: "ali3.bagheri.software@gmail.com",
		Phone: "1233454321",
	})
}