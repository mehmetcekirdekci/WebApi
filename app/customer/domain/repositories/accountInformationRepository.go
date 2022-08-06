package repositories

import (
	"errors"

	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
	"gorm.io/gorm"
)

type AccountInformationRepository interface {
	InsertAccountInformation(accountInformation *types.AccountInformation) error
}

type (
	accountInformationRepository struct {
		db *gorm.DB
	}
)

func NewAccountInformationRepository(database *gorm.DB) AccountInformationRepository {
	return &accountInformationRepository{
		db: database,
	}
}

func (receiver *accountInformationRepository) InsertAccountInformation(accountInformation *types.AccountInformation) error {
	err := receiver.db.Table(types.AccountInformationsTable).Create(&accountInformation).Error
	if err != nil {
		return errors.New("Account information can not be inserted.")
	}
	return nil
}