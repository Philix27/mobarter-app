package announcement

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(data createAnnouncementDto) error
	Update(data updateAnnouncementDto) error
	Delete(dataId int) error
	FindById(dataId int) (announcement, error)
	FindAll() ([]announcement, error)
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
func (r *Repository) Create(data createAnnouncementDto) error {

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

// Update implements iRepository.
func (r *Repository) Update(data updateAnnouncementDto) error {
	result := r.Db.Where(
		"id = ?", data.Id,
	).Updates(data)

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil
}

// FindAll implements iRepository.
func (r *Repository) FindAll() ([]announcement, error) {
	var announceList []announcement
	result := r.Db.Limit(10).Find(&announceList)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ALL:", result.Error)
		return nil, result.Error
	}
	r.logger.Info("READ ALL")
	return announceList, nil
}

// FindById implements iRepository.
func (r *Repository) FindById(dataId int) (announcement, error) {
	announcementItem := announcement{}

	result := r.Db.First(&announcementItem, dataId)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error, dataId)
		return announcement{}, result.Error
	} else {
		r.logger.Error("RETRIEVE FIND_ONE:", announcementItem)
		return announcementItem, nil
	}

}

// FindById implements iRepository.
func (r *Repository) FindByProperties(dataId int) (announcement, error) {
	var announcementItem announcement
	result := r.Db.Where(
		"id = ?", dataId,
	).First(&announcementItem)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error)
		return announcement{}, result.Error
	} else {
		return announcementItem, nil
	}

}

// Delete implements iRepository.
func (r *Repository) Delete(dataId int) error {
	result := r.Db.Where("id = ?", dataId).Delete(new(announcement))

	if result.Error != nil {
		return result.Error
	}
	return nil
}
