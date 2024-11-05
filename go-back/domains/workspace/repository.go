package workspace

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(WorkspaceCreateInput) error
	Update(WorkspaceUpdateInput) error
	Delete(int) error
	FindByUserId(int) (Workspace, error)
	FindAllByUserId(WorkspaceGetAllInput) ([]Workspace, error)
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
func (r *Repository) Create(data WorkspaceCreateInput) error {
	// fmt.Println(data)
	if result := r.Db.Create(&database.Workspace{
		Name:        data.Name,
		Description: data.Description,
		OwnerId:     data.OwnerId,
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
func (r *Repository) Update(data WorkspaceUpdateInput) error {
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
func (r *Repository) FindAllByUserId(data WorkspaceGetAllInput) ([]Workspace, error) {
	var list []Workspace
	result := r.Db.Where("owner_id =? ", data.UserId).Limit(data.Limit).Find(&list)

	if result.Error != nil {
		r.logger.Error("CANNOT_FIND_ALL:", result.Error)
		return nil, result.Error
	}
	r.logger.Info("READ_ALL_SUCCESS")
	return list, nil
}

// FindById implements iRepository.
func (r *Repository) FindByUserId(id int) (Workspace, error) {
	object := Workspace{}

	result := r.Db.Where("owner_id = ?", id).First(&object, id)

	if result.Error != nil {
		r.logger.Error("CANNOT_FIND_ONE:", result.Error)
		return Workspace{}, result.Error
	} else {
		r.logger.Error("RETRIEVE_FIND_ONE:", object)
		return object, nil
	}

}

// Delete implements iRepository.
func (r *Repository) Delete(dataId int) error {
	result := r.Db.Where("id = ?", dataId).Delete(&database.Workspace{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
