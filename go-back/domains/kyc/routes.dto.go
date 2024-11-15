package waitlist

type CreateInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateResponse struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

type GetOneInput struct {
	Id int `json:"id"`
}
type GetOneResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type GetAllInput struct {
	Limit int `json:"limit"`
}

type GetAllResponse struct {
	Data []KycCredential `json:"data"`
}

type KycCredential struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
