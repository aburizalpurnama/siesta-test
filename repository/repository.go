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
		GetLendingById(id uint) (model.Lending, error)
		UpdateLendingStatus(data model.Lending) error
		SelectRepaymentsByAccountId(id uint) ([]model.Repayment, error)
	}

	MbtRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewMbtRepo(db *gorm.DB) MbtRepository {
	return &MbtRepositoryImpl{db}
}

func (r *MbtRepositoryImpl) SelectRepaymentsByAccountId(id uint) ([]model.Repayment, error) {
	sql := `SELECT r.* FROM repayments r JOIN lendings l ON r.lending_id = l.id WHERE l.account_id = ? ORDER BY date_part('year', r.due_date), date_part('month', r.due_date) ASC;`
	repayments := []model.Repayment{}
	err := r.db.Raw(sql, id).Scan(&repayments).Error
	if err != nil {
		log.Info(err)
	}

	return repayments, err
}

func (r *MbtRepositoryImpl) GetLendingById(id uint) (model.Lending, error) {
	data := model.Lending{}
	err := r.db.First(&data, id).Error
	if err != nil {
		log.Info(err)
	}

	return data, err
}

func (r *MbtRepositoryImpl) UpdateLendingStatus(data model.Lending) error {
	err := r.db.Model(&data).Update("status", data.Status).Error
	if err != nil {
		log.Info(err)
	}

	return err
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
