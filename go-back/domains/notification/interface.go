package notification

type iService interface {
	SendEmailOtp()
	SendPhoneOtp()
	SendEmail(email string, msg string)
	SendPhoneMsg()
}
