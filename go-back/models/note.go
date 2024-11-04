package models

type Note struct {
	Id      int    `gorm:"type:int;primary_key"`
	Content string `gorm:"not null" json:"content, omitempty"`
}

type NoteResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type CreateNoteRequest struct {
	Content string `validate:"required,min=2,max=100" json:"content"`
}
type UpdateNoteRequest struct {
	Id string `validate:"required`
	Content string `validate:"required,min=2,max=100" json:"content"`
}
