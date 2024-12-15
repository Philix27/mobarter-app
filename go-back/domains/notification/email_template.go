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
		return "One Time Password"

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}

func getMailPlainText(t MailTemplates, msg string) string {
	switch t {
	case OTP:
		return `
Hello,

You recently requested for a one time password from Mobarter.

Please use the one-time password (OTP) below to confirm it's really you.
OTP code will expire in 10 mins.

Remember, never share your OTP with anyone. Our team will never ask for your 
password or OTP, so please keep this information private. 
Stay cautious and protect your account.`

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}

func getMailHtmlContent(t MailTemplates, msg string) string {
	switch t {
	case OTP:
		return OtpHtmlEmailTemplate(msg)
		// return "Your one time password is: <br /> <strong>" + msg + "</strong>"

	case NewAccount:
		return "Congratulations, your account has been created/"
	}
	return "New Mail"
}
