package workspaceEpics

import (
	"mobarter/database"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

var ModuleName = "WORKSPACE_EPIC"

func Setup(router fiber.Router, db *gorm.DB, logger *slog.Logger) {

	repo := NewRepository(
		db.Model(&database.WorkspaceEpic{}).Debug(),
		logger, ModuleName+"_REPOSITORY")

	handler := NewRoutes(repo, logger, "_ROUTES")

	router.Route("/workspace_epic", handler.manager)

}
