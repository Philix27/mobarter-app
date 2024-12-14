package notification

type MailTemplates int

const (
	OTP MailTemplates = iota + 1
	NewAccount
	Congratulations
)

func IsValidMailTemplate(t MailTemplates) bool {
	switch t {
	case OTP, NewAccount, Congratulations:
		return true
	}
	return false
}

func getMailSubject(t MailTemplates) string {
	switch t {
	case OTP:
		return "One time password"

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}

func getMailPlainText(t MailTemplates, msg string) string {
	switch t {
	case OTP:
		return "Your one time password is " + msg + ". It will expire in 10 mins"

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}
func getMailHtmlContent(t MailTemplates, msg string) string {
	switch t {
	case OTP:
		return "<strong>" + msg + "</strong>"

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}
