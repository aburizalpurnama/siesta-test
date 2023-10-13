package repository

import (
	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type (
	MbtRepository interface {
		GetLendingTypeById(id uint) (model.LendingType, error)
		GetAccountById(id uint) (model.Account, error)
		CreateRepayment(data model.Repayment) (uint, error)
		CreateLending(data model.Lending) (uint, error)
	}

	MbtRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewMbtRepo(db *gorm.DB) MbtRepository {
	return &MbtRepositoryImpl{db}
}

func (r *MbtRepositoryImpl) GetLendingTypeById(id uint) (model.LendingType, error) {
	data := model.LendingType{}
	err := r.db.First(&data, id).Error
	if err != nil {
		log.Info(err)
	}

	return data, err
}

func (r *MbtRepositoryImpl) GetAccountById(id uint) (model.Account, error) {
	data := model.Account{}
	err := r.db.First(&data, id).Error
	if err != nil {
		log.Info(err)
	}

	return data, err
}

func (r *MbtRepositoryImpl) CreateLending(data model.Lending) (uint, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		log.Info(err)
	}

	return data.ID, err
}

func (r *MbtRepositoryImpl) CreateRepayment(data model.Repayment) (uint, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		log.Info(err)
	}

	return data.ID, err
}
