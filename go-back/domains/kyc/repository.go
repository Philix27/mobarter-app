package waitlist

import (
	"mobarter/database"

	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type iRepository interface {
	Create(data CreateInput) error
	FindById(dataId int) (KycCredential, error)
	Bvn(bvn string, userId string) error
	Nin(nin string, userId string) error
	Phone(phoneNo string, userId string) error
	Email(email string, userId string) error
	HomeAddress(address string, userId string) error
	NextOfKin(kinName string, kinPhone string, userId string) error
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

func (r *Repository) FindById(dataId int) (KycCredential, error) {
	waitlistItem := KycCredential{}

	result := r.Db.First(&waitlistItem, dataId)

	if result.Error != nil {
		r.logger.Error("CANNOT FIND_ONE:", result.Error, dataId)
		return KycCredential{}, result.Error
	} else {
		r.logger.Error("RETRIEVE FIND_ONE:", waitlistItem)
		return waitlistItem, nil
	}

}
func (r *Repository) Bvn(bvn string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		Bvn: bvn,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}
func (r *Repository) Nin(nin string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		Nin: nin,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}
func (r *Repository) Phone(phoneNo string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		Phone: phoneNo,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}
func (r *Repository) Email(email string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		Email: email,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}
func (r *Repository) HomeAddress(address string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		HomeAddress: address,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}

func (r *Repository) NextOfKin(kinName string, kinPhone string, userId string) error {
	result := r.Db.Where(
		"id = ?", userId,
	).Updates(&database.KycCredential{
		NextOfKinName:  kinName,
		NextOfKinPhone: kinPhone,
	})

	if result.Error != nil {
		r.logger.Error("CANNOT UPDATE")
		return result.Error
	}
	r.logger.Info("UPDATE announcement successful")
	return nil

}
