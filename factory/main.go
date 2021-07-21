package main

import "fmt"

//SMS, Email

type INotificationFactory interface {
	SendNotification(message string)
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type Notification struct {
	source string
}

//SMS Structs
type SMSNotification struct {
	Notification
}

func (SMSNotification) SendNotification(message string) {
	fmt.Println("SMS Message: ", message, " was sent")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twillio"
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

//Email structs
type EmailNotification struct {
	Notification
}

func (EmailNotification) SendNotification(message string) {
	fmt.Println("Email Message: ", message, " was sent")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "SMPT:Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Exchange server"
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Notification Type " + notificationType + " doesn't exist")
}
func sendNotification(f INotificationFactory, message string) {
	f.SendNotification(message)
}
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory, "This is an SMS Text message")
	getMethod(smsFactory)
	sendNotification(emailFactory, "This is an Email message")
	getMethod(emailFactory)

}
