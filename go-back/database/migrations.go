package database

import (
	"gorm.io/gorm"
)

// The order of which you run the migrations matters
// User first, workspace next
// Tables with More dependencies stays on top
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		Workspace{},
		WorkspaceEpic{},
		WorkspaceStory{},
		Task{},
		Board{},
		BoardStage{},
		User{},
		UserAuth{},
		Announcement{},
		Note{},
		Canvas{},
		Document{},
		TaskTags{},
		Schedule{},
	)

	if err != nil {
		println("Could not run migrations, %v", err.Error())
	}
}
