package main

import "fmt"

type EmailProvider struct {
}

type SMSProvider struct {
}

func (ep *EmailProvider) Send(recipient string, message string) {
	fmt.Printf("PROVIDER: Enviando E-MAIL para %s: '%s'\n", recipient, message)

}

func (ep *SMSProvider) Send(recipient string, message string) {
	fmt.Printf("PROVIDER: Enviando SMS para %s: '%s'\n", recipient, message)

}

type NotificationProvider interface {
	Send(recipient string, message string)
}

type NotificationService struct {
	provider NotificationProvider
}

func (ns *NotificationService) SendNotification(recipient string, message string) {
	ns.provider.Send(recipient, message)
}

func main() {
	provider := EmailProvider{}

	service := NotificationService{provider: &provider}

	service.SendNotification("exemplo@email.com", "Olá email!")

	SMSProvider := SMSProvider{}

	SMSService := NotificationService{provider: &SMSProvider}

	SMSService.SendNotification("exemplo@sms.com", "Olá SMS!")
}
