package workspaceEpics

// * Create
type WorkspaceEpicCreateInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}

type WorkspaceEpicCreateResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}

type WorkspaceEpicCreatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceEpicCreateQueryParams struct {
}

// * Update
type WorkspaceEpicUpdateInput struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type WorkspaceEpicUpdateResponse struct {
	Msg string `json:"msg"`
}

type WorkspaceEpicUpdatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceEpicUpdateQueryParams struct {
}

// * Delete
type WorkspaceEpicDeleteResponse struct {
	Msg string `json:"msg"`
}
type WorkspaceEpicDeletePathParams struct {
	Id string `json:"id"`
}
type WorkspaceEpicDeleteQueryParams struct {
}

// * Read one
type WorkspaceEpicGetOneResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}

type WorkspaceEpicGetOneInput struct {
}
type WorkspaceEpicGetOnePathParams struct {
}
type WorkspaceEpicGetOneQueryParams struct {
}

// * Read all

type WorkspaceEpicGetAllResponse struct {
	Data []WorkspaceEpic `json:"data"`
}

type WorkspaceEpicGetAllInput struct {
}

type WorkspaceEpicGetAllPathParams struct {
}
type WorkspaceEpicGetAllQueryParams struct {
	Workspace_id string `json:"id"`
	Limit        string `json:"limit"`
}

type WorkspaceEpic struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WorkspaceId uint   `json:"workspaceId"`
}
