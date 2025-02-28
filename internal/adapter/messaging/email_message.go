package messaging

import (
	"crypto/tls"
	"latihan-compro/config"

	"github.com/go-mail/mail"
	"github.com/labstack/gommon/log"
)

type EmailMessagingInterface interface {
	SendEmailAppointment(attach *string, from, subject, body string) error
}

type emailAttributes struct {
	username string
	password string
	host     string
	port     int
	isTLS    bool
	receiver string
}

// SendEmailAppointment implements EmailMessagingInterface.
func (e *emailAttributes) SendEmailAppointment(attach *string, from, subject string, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", e.receiver)

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if attach != nil {
		m.Attach(*attach)
	}

	d := mail.NewDialer(e.host, e.port, e.username, e.password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if err := d.DialAndSend(m); err != nil {
		log.Errorf("error sending mail: %v", err)
		return err
	}
	return nil
}

func NewEmailMessaging(cfg *config.Config) EmailMessagingInterface {
	return &emailAttributes{
		username: cfg.Email.Username,
		password: cfg.Email.Password,
		host:     cfg.Email.Host,
		port:     cfg.Email.Port,
		isTLS:    cfg.Email.IsTLS,
		receiver: cfg.Email.Reciever,
	}
}
