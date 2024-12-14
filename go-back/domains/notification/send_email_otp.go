package notification

func SendEmailOtp(email string, otp string) error {
	err := SendMail(OTP, email, otp)

	return err
}

func SendPhoneOtp(phone string, countryCode string) error {
	// todo
	return nil
}
