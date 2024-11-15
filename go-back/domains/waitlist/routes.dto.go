package waitlist

// * Create
type CreateInput struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	WorkspaceId uint   `json:"workspaceId"`
}

type CreateResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

// * Read one
type GetOneInput struct {
	Id int `json:"id"`
}
type GetOneResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

// * Read all
type GetAllInput struct {
	Limit int `json:"limit"`
}
type GetAllResponse struct {
	Data []waitlist `json:"data"`
}

type waitlist struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type createDto struct {
	Title       string `validate:"required, min=1, max=10" json:"name" `
	Subtitle    string `validate:"required"`
	WorkspaceId uint
}

type updateDto struct {
	Id       int `validate:"required"`
	Title    string
	Subtitle string
}
