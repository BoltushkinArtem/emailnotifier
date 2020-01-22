package main

import (
	"fmt"
	email "github.com/BoltushkinArtem/emailnotifier/backend/app/internal"
)

func main() {
	notifier := email.Notifier{
		Identity: "",
		Username: "user@example.com",
		Password: "password",
		Host:     "mail.example.com",
		Port:     "25",
	}

	message := email.Message{
		From:    "Test From",
		To:      []string{"userto@example.com"},
		Subject: "Subject",
		Message: "Message",
	}
	if err := notifier.SendMessage(message); err != nil {
		fmt.Println(err)
	}
}
