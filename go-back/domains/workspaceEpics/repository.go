package workspaceEpics

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(WorkspaceEpicCreateInput) error
	Update(WorkspaceEpicUpdateInput) error
	Delete(int) error
	FindById(int) (WorkspaceEpic, error)
	FindWorkspaceId(workspaceId int, limit int) ([]WorkspaceEpic, error)
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
func (r *Repository) Create(data WorkspaceEpicCreateInput) error {
	// fmt.Println(data)
	if result := r.Db.Create(&database.WorkspaceEpic{
		Title:       data.Title,
		Description: data.Description,
		WorkspaceID: data.WorkspaceId,
	}); result.Error != nil {
		println(result.Error)
		r.logger.Error(
			"FAILED_TO_CREATE",
			"errMsg",
			result.Error,
		)
		return result.Error
	}
	r.logger.Info("CREATE_SUCCESSFUL")
	return nil
}

// Update implements iRepository.
func (r *Repository) Update(data WorkspaceEpicUpdateInput) error {
	result := r.Db.Where(
		"id = ?", data.Id,
	).Updates(data)

	if result.Error != nil {
		r.logger.Error("CANNOT_UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE_SUCCESSFUL")
	return nil
}

// FindAll implements iRepository.
func (r *Repository) FindWorkspaceId(workspaceId int, limit int) ([]WorkspaceEpic, error) {
	var list []WorkspaceEpic
	result := r.Db.Where("workspace_id =? ", workspaceId).Limit(limit).Find(&list)

	if result.Error != nil {
		r.logger.Error("CANNOT_FIND_ALL:", result.Error)
		return nil, result.Error
	}
	r.logger.Info("READ_ALL_SUCCESS")
	return list, nil
}

// FindById implements iRepository.
func (r *Repository) FindById(id int) (WorkspaceEpic, error) {
	object := WorkspaceEpic{}

	result := r.Db.Where("id = ?", id).First(&object, id)

	if result.Error != nil {
		r.logger.Error("FindById:", result.Error)
		return WorkspaceEpic{}, result.Error
	} else {
		r.logger.Error("RETRIEVE_FIND_ONE:", object)
		return object, nil
	}

}

// Delete implements iRepository.
func (r *Repository) Delete(dataId int) error {
	result := r.Db.Where("id = ?", dataId).Delete(&database.WorkspaceEpic{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
