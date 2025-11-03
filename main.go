package main

import (
	"fmt"
	"time"
)

type Destinatary struct {
	Email    string
	Telefone string
}

type EmailProvider struct {
}

type SMSProvider struct {
}

func (ep *EmailProvider) Send(dest Destinatary, message string) error {
	if dest.Email == "" {
		return fmt.Errorf("E-mail do destinatário não fornecido")
	}
	fmt.Printf("PROVIDER: Enviando E-MAIL para %s: '%s'\n", dest.Email, message)
	return nil
}

func (e *EmailProvider) Schedule(dest Destinatary, message string, sendAt time.Time) error {
	fmt.Println("Agendando E-mail...")
	return nil
}

func (ep *SMSProvider) Send(dest Destinatary, message string) error {
	if dest.Telefone == "" {
		return fmt.Errorf("Telefone do destinatário não fornecido")
	}
	fmt.Printf("PROVIDER: Enviando SMS para %s: '%s'\n", dest.Telefone, message)
	return nil
}

type Sender interface {
	Send(dest Destinatary, message string) error
}

type Scheduler interface {
	Schedule(dest Destinatary, message string, sendAt time.Time) error
}

type NotificationService struct {
	provider Sender
}

func (ns *NotificationService) SendNotification(dest Destinatary, message string) error {
	return ns.provider.Send(dest, message)
}
func (ns *NotificationService) ScheduleNotification(dest Destinatary, message string, sendAt time.Time) error {

	scheduler, ok := ns.provider.(Scheduler)

	if !ok {
		return fmt.Errorf("este provedor não suporta agendamento")
	}

	return scheduler.Schedule(dest, message, sendAt)
}

func main() {
	emailSvc := NotificationService{provider: &EmailProvider{}}
	err1 := emailSvc.ScheduleNotification(Destinatary{}, "msg", time.Now())

	smsSvc := NotificationService{provider: &SMSProvider{}}
	err2 := smsSvc.ScheduleNotification(Destinatary{}, "msg", time.Now())

	fmt.Println("Resultado do E-mail:", err1)
	fmt.Println("Resultado do SMS:", err2)

}
