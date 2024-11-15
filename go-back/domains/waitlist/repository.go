package waitlist

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(data CreateInput) error
	FindById(dataId int) (Waitlist, error)
	FindAll() ([]Waitlist, error)
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

func (r *Repository) Create(data CreateInput) error {

	result := r.Db.Create(&database.Waitlist{
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
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

func (r *Repository) FindAll() ([]Waitlist, error) {
	var list []Waitlist
	result := r.Db.Limit(10).Find(&list)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ALL:", result.Error)
		return nil, result.Error
	}
	r.logger.Info("READ ALL")
	return list, nil
}

func (r *Repository) FindById(dataId int) (Waitlist, error) {
	waitlistItem := Waitlist{}

	result := r.Db.First(&waitlistItem, dataId)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error, dataId)
		return Waitlist{}, result.Error
	} else {
		r.logger.Error("RETRIEVE FIND_ONE:", waitlistItem)
		return waitlistItem, nil
	}

}
