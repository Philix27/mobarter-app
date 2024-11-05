package workspaceStory

// * Create
type WorkspaceStoryCreateInput struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	WorkspaceEpicId uint   `json:"workspaceEpicId"`
}

type WorkspaceStoryCreateResponse struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	WorkspaceEpicId uint   `json:"workspaceEpicId"`
}

type WorkspaceStoryCreatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceStoryCreateQueryParams struct {
}

// * Update
type WorkspaceStoryUpdateInput struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type WorkspaceStoryUpdatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceStoryUpdateQueryParams struct {
}
type WorkspaceStoryUpdateResponse struct {
	Msg string `json:"msg"`
}

// * Delete
type WorkspaceStoryDeleteResponse struct {
	Msg string `json:"msg"`
}
type WorkspaceStoryDeletePathParams struct {
	Id string `json:"id"`
}
type WorkspaceStoryDeleteQueryParams struct {
}

// * Read one
type WorkspaceStoryGetOneResponse struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	WorkspaceEpicId uint   `json:"workspaceEpicId"`
}
type WorkspaceStoryGetOneInput struct {
}

type WorkspaceStoryGetOnePathParams struct {
}
type WorkspaceStoryGetOneQueryParams struct {
}
// * Read all
type WorkspaceStoryGetAllResponse struct {
	Data []WorkspaceStory `json:"data"`
}

type WorkspaceStoryGetAllInput struct {
}
type WorkspaceStoryGetAllPathParams struct {
}
type WorkspaceStoryGetAllQueryParams struct {
	Epic_id string
	Limit   string
}

type WorkspaceStory struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	WorkspaceEpicId uint   `json:"workspaceEpicId"`
}
