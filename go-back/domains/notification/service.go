package notification

type Service struct {
	SendGrid string
	Twilo    string
	Termii   string
}

// sendEmail implements iService.
func (s *Service) SendEmail(email string, msg string) {
	panic("unimplemented")
}

// sendEmailOtp implements iService.
func (*Service) SendEmailOtp() {
	panic("unimplemented")
}

// sendPhoneMsg implements iService.
func (*Service) SendPhoneMsg() {
	panic("unimplemented")
}

// sendPhoneOtp implements iService.
func (*Service) SendPhoneOtp() {
	panic("unimplemented")
}

func InitializeNotification() iService {
	return &Service{}
}
