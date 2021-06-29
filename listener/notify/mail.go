package notify

import (
	"errors"
	"fmt"
	"github.com/go-gomail/gomail"
	"strings"
	"supervisor-event-listener/event"
)

type Mail struct{}

func (mail *Mail) Send(message event.Message) error {
	body := message.String()
	body = strings.Replace(body, "\n", "<br>", -1)
	gomailMessage := gomail.NewMessage()
	gomailMessage.SetHeader("From", Conf.MailServer.User)
	gomailMessage.SetHeader("To", Conf.MailUser.Email...)
	gomailMessage.SetHeader("Subject", Conf.MailServer.Subject)
	gomailMessage.SetBody("text/html", body)
	mailer := gomail.NewDialer(
		Conf.MailServer.Host,
		Conf.MailServer.Port,
		Conf.MailServer.User,
		Conf.MailServer.Password,
	)
	err := mailer.DialAndSend(gomailMessage)
	if err == nil {
		return nil
	}
	errorMessage := fmt.Sprintf("邮件发送失败#%s", err.Error())

	return errors.New(errorMessage)
}
