package notification

import (
	"fmt"

	"github.com/resend/resend-go/v2"
)

func SendMail(tem MailTemplates, receiverEmail string, msg string) error {

	// client := resend.NewClient(config.EnvValues.ResendKey)
	client := resend.NewClient("re_Y7qoNYXo_4a6GG1tN9t15Bx4CqbGP3eTk")

	params := &resend.SendEmailRequest{
		To:   []string{receiverEmail},
		From: "Mobarter <notifications@mobarter.com>",
		// Text:    getMailPlainText(tem, msg),
		Html:    getMailHtmlContent(tem, msg),
		Subject: getMailSubject(tem),
		// Cc:      []string{"cc@example.com"},
		// Bcc:     []string{"cc@example.com"},
		// ReplyTo: "replyto@example.com",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(sent.Id)

	return err
}
