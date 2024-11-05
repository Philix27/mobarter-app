package announcement

import (
	"mobarter/database"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

var ModuleName = "ANNOUNCEMENT"

func Setup(router fiber.Router, db *gorm.DB, logger *slog.Logger) {

	repo := NewRepository(
		db.Model(&database.Announcement{}).Debug(),
		logger, ModuleName+"_REPOSITORY")

	handler := NewRoutes(repo, logger, ModuleName+"_ROUTES")

	router.Route("/announcement", handler.manager)

}
