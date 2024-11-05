package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Password    string
	WorkspaceID uint
	UserAuth    uint
	ScheduleID  []Schedule
}
type UserAuth struct {
	gorm.Model
	Email          string
	Password       string
	AccessToken    string
	RefreshToken   string
	SlackToken     string
	LinkedInToken  string
	TwitterToken   string
	WhatsappToken  string
	InstagramToken string
	DiscordToken   string
	UserID         uint
}

// / * Workspace
type Workspace struct {
	gorm.Model
	Name          string
	Description   string
	OwnerId       uint
	Notes         []Document
	Announcements []Announcement
	Epics         []WorkspaceEpic
	Members       []User
	Boards        []Board
}

type WorkspaceEpic struct {
	gorm.Model
	Title            string
	Description      string
	WorkspaceID      uint
	WorkspaceStories []WorkspaceStory
}

type WorkspaceStory struct {
	gorm.Model
	Title       string
	Description string
	Tasks       []Task
	WorkspaceEpicID  uint
}

type Board struct {
	gorm.Model
	Title       string
	Description string
	WorkspaceID uint
	BoardStages []BoardStage
}
type BoardStage struct {
	gorm.Model
	Title       string
	Description string
	Color       string
	BoardID     uint
	Tasks       []Task
}

type Task struct {
	gorm.Model
	Name              string
	Description       string
	GithubIssueNumber string
	GithubIssueLink   string
	AssignedTo        string
	WorkspaceStoryID  uint
	TaskEpic          uint
	BoardStageID      uint
}

type Announcement struct {
	gorm.Model
	Title       string
	Subtitle    string
	WorkspaceID uint
}

type TaskTags struct {
	gorm.Model
	Name        string
	Description string
	ColorCode   string
}

// / * Schedule
type Schedule struct {
	gorm.Model
	Title       string
	Msg         string
	ImgPath     string
	Type        string
	IsCompleted bool
	IsCanceled  bool
	UserID      uint
}

/// * Document

type Document struct {
	gorm.Model
	Title       string
	Body        string
	WorkspaceID uint
	Note        []Note
	Canvas      []Canvas
}
type Note struct {
	gorm.Model
	Title       string
	Body        string
	Archived    bool
	HasChild    bool
	WorkspaceID uint
	DocumentID  uint
}
type Canvas struct {
	gorm.Model
	Title       string
	Body        string
	Archived    bool
	HasChild    bool
	WorkspaceID uint
	DocumentID  uint
}
