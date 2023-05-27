package notification

import (
	"io/ioutil"

	"github.com/piovani/wallet/infra/config"
	"gopkg.in/gomail.v2"
)

func (n *Notification) GetMsgEmail() *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "from@example.com")
	msg.SetHeader("To", "to@example.com")
	msg.SetHeader("subject", "teste")

	text, _ := ioutil.ReadFile("./infra/notification/email.html")
	msg.SetBody("text/html", string(text))

	return msg
}

func (n *Notification) SendEmail(msg *gomail.Message) error {
	dialer := gomail.NewDialer(
		config.Env.EmailHost,
		int(config.Env.EmailPort),
		config.Env.EmailUser,
		config.Env.EmailPassword,
	)

	return dialer.DialAndSend(msg)
}
