package workspace

// * Create
type WorkspaceCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     uint   `json:"ownerId"`
}

type WorkspaceCreateResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WorkspaceCreatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceCreateQueryParams struct {
}

// * Update
type WorkspaceUpdateInput struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type WorkspaceUpdateResponse struct {
	Msg string `json:"msg"`
}
type WorkspaceUpdatePathParams struct {
	Id string `json:"id"`
}
type WorkspaceUpdateQueryParams struct {
}

// * Delete
type WorkspaceDeleteResponse struct {
	Msg string `json:"msg"`
}
type WorkspaceDeleteInput struct {
	Id string `json:"id"`
}

type WorkspaceDeletePathParams struct {
	Id string `json:"id"`
}
type WorkspaceDeleteQueryParams struct {
}

// * Read one
type WorkspaceGetOneResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type WorkspaceGetOneInput struct {
}
type WorkspaceGetOnePathParams struct {
}
type WorkspaceGetOneQueryParams struct {
}

// * Read all
type WorkspaceGetAllInput struct {
	Limit  int `json:"limit"`
	UserId int `json:"userId"`
}
type WorkspaceGetAllResponse struct {
	Data []Workspace `json:"data"`
}
type WorkspaceGetAllPathParams struct {
	Id string `json:"id"`
}
type WorkspaceGetAllQueryParams struct {
}

type Workspace struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
