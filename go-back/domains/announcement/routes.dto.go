package announcement

// * Create
type AnnouncementCreateInput struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	WorkspaceId uint   `json:"workspaceId"`
}

type AnnouncementCreateResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type AnnouncementCreatePathParams struct {
	Id string `json:"id"`
}
type AnnouncementCreateQueryParams struct {
}

// * Update
type AnnouncementUpdateInput struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
type AnnouncementUpdateResponse struct {
	Msg string `json:"msg"`
}

type AnnouncementUpdatePathParams struct {
	Id string `json:"id"`
}
type AnnouncementUpdateQueryParams struct {
}

// * Delete

type AnnouncementDeleteInput struct {
	Id int `json:"id"`
}
type AnnouncementDeleteResponse struct {
	Msg string `json:"msg"`
}

type AnnouncementDeletePathParams struct {
	Id string `json:"id"`
}
type AnnouncementDeleteQueryParams struct {
}

// * Read one
type AnnouncementGetOneInput struct {
	Id int `json:"id"`
}
type AnnouncementGetOneResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type AnnouncementGetOnePathParams struct {
	Id string `json:"id"`
}
type AnnouncementGetOneQueryParams struct {
}

// * Read all
type AnnouncementGetAllInput struct {
	Limit int `json:"limit"`
}
type AnnouncementGetAllResponse struct {
	Data []announcement `json:"data"`
}

type AnnouncementGetAllPathParams struct {
	Id string `json:"id"`
}
type AnnouncementGetAllQueryParams struct {
}

type announcement struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type createAnnouncementDto struct {
	Title       string `validate:"required, min=1, max=10" json:"name" `
	Subtitle    string `validate:"required"`
	WorkspaceId uint
}

type updateAnnouncementDto struct {
	Id       int `validate:"required"`
	Title    string
	Subtitle string
}
