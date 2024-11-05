package board

// * Create
type BoardCreateInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}
type BoardCreateResponse struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	WorkspaceEpicId uint   `json:"workspaceEpicId"`
}
type BoardCreatePathParams struct {
	Id string `json:"id"`
}
type BoardCreateQueryParams struct {
}

// * Update
type BoardUpdateInput struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type BoardUpdatePathParams struct {
	Id string `json:"id"`
}
type BoardUpdateQueryParams struct {
}
type BoardUpdateResponse struct {
	Msg string `json:"msg"`
}

// * Delete
type BoardDeleteResponse struct {
	Msg string `json:"msg"`
}
type BoardDeletePathParams struct {
	Id string `json:"id"`
}
type WorkspaceStoryDeleteInput struct {
}
type BoardDeleteQueryParams struct {
}

// * Read one
type BoardGetOneResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}
type BoardGetOnePathParams struct {
}
type BoardGetOneQueryParams struct {
}

// * Read all
type BoardGetAllResponse struct {
	Data []Board `json:"data"`
}

type BoardGetAllInput struct {
}
type BoardGetAllPathParams struct {
	Id string `json:"id"`
}
type BoardGetAllQueryParams struct {
	Workspace_id string
	Limit        string
}

type Board struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}
