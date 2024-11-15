package waitlist

import (
	"mobarter/database"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

var ModuleName = "WAIT_LIST"

func Setup(router fiber.Router, db *gorm.DB, logger *slog.Logger) {

	repo := NewRepository(
		db.Model(&database.Waitlist{}).Debug(),
		logger, ModuleName+"_REPOSITORY")

	handler := NewRoutes(repo, logger, ModuleName+"_ROUTES")

	router.Route("/waitlist", handler.manager)

}
