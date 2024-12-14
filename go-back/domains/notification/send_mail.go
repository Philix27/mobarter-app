package notification

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(tem MailTemplates, receiverEmail string, msg string) error {
	from := mail.NewEmail("app.mobarter.com", "app@mobarter.com")
	subject := getMailSubject(tem)
	to := mail.NewEmail("user", receiverEmail)
	plainTextContent := getMailPlainText(tem, msg)
	htmlContent := getMailHtmlContent(tem, msg)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return err
}
