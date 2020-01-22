package email

import (
	"net/smtp"
	"strings"
)

type Notifier struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     string
}

type Message struct {
	From    string
	To      []string
	Subject string
	Message string
}

func (n *Notifier) SendMessage(m Message) error {
	return n.sendMail(m.From, m.To, m.Subject, m.Message)
}

func (n *Notifier) sendMail(from string, to []string, subject string, message string) error {
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n" + message

	return smtp.SendMail(
		n.Host+":"+n.Port,
		smtp.PlainAuth(
			n.Identity,
			n.Username,
			n.Password,
			n.Host,
		),
		n.Username,
		to,
		[]byte(msg))
}
