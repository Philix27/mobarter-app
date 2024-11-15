package waitlist

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(data createDto) error
	FindById(dataId int) (waitlist, error)
	FindAll() ([]waitlist, error)
}

type Repository struct {
	Db          *gorm.DB
	logger      *slog.Logger
	logGroupKey string
}

func NewRepository(db *gorm.DB, logger *slog.Logger, logGroupKey string) iRepository {
	return &Repository{
		Db:          db,
		logger:      logger,
		logGroupKey: logGroupKey,
	}
}

// Create implements iRepository.
func (r *Repository) Create(data createDto) error {

	result := r.Db.Create(&database.Announcement{
		Title:       data.Title,
		Subtitle:    data.Subtitle,
		WorkspaceID: data.WorkspaceId,
	})

	if result.Error != nil {
		println(result.Error)
		r.logger.WithGroup(r.logGroupKey).Error(
			"FAILED_TO_CREATE",
			"errMsg",
			result.Error,
		)
		return result.Error
	}
	r.logger.WithGroup(r.logGroupKey).Info("CREATE_SUCCESSFUL")
	return nil
}

// FindAll implements iRepository.
func (r *Repository) FindAll() ([]waitlist, error) {
	var announceList []waitlist
	result := r.Db.Limit(10).Find(&announceList)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ALL:", result.Error)
		return nil, result.Error
	}
	r.logger.Info("READ ALL")
	return announceList, nil
}

// FindById implements iRepository.
func (r *Repository) FindById(dataId int) (waitlist, error) {
	waitlistItem := waitlist{}

	result := r.Db.First(&waitlistItem, dataId)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error, dataId)
		return waitlist{}, result.Error
	} else {
		r.logger.Error("RETRIEVE FIND_ONE:", waitlistItem)
		return waitlistItem, nil
	}

}

// FindById implements iRepository.
func (r *Repository) FindByProperties(dataId int) (waitlist, error) {
	var waitlistItem waitlist
	result := r.Db.Where(
		"id = ?", dataId,
	).First(&waitlistItem)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error)
		return waitlist{}, result.Error
	} else {
		return waitlistItem, nil
	}

}
