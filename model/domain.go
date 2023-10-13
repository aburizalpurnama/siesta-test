package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Repayment struct {
		gorm.Model
		DueDate      time.Time `gorm:"not null"`
		AdminFee     int64     `gorm:"not null"`
		StampDutyFee int64     `gorm:"not null"`
		Interest     int64     `gorm:"not null"`
		Principal    int64     `gorm:"not null"`
		Bill         int64     `gorm:"not null"`
		Outstanding  int64     `gorm:"not null"`
		IsPaid       bool      `gorm:"not null"`
		LendingID    uint      `gorm:"not null, index"`
	}

	Lending struct {
		gorm.Model
		Date              time.Time `gorm:"not null"`
		Amount            int64     `gorm:"not null"`
		Tenor             int       `gorm:"not null"`
		Principal         int64     `gorm:"not null"`
		PrincipalPerMonth int64     `gorm:"not null"`
		MonthlyInterest   int64     `gorm:"not null"`
		TotalInterest     int64     `gorm:"not null"`
		AdminFee          int64     `gorm:"not null"`
		LendingPurpose    string    `gorm:"not null"`
		WithdrawalType    string    `gorm:"not null"`
		LinkPurchase      string
		Status            string `gorm:"not null"`
		AccountID         uint   `gorm:"not null, index"`
		LendingTypeID     uint   `gorm:"not null, index"`

		Repayments []Repayment
	}

	Account struct {
		gorm.Model
		Name           string `gorm:"not null"`
		Limit          int64  `gorm:"not null"`
		LimitRemaining int64  `gorm:"not null"`

		Lendings []Lending
	}

	LendingType struct {
		gorm.Model
		Name             string  `gorm:"not null"`
		MonthlyInterest  float64 `gorm:"not null"`
		AdminFee         float64 `gorm:"not null"`
		StampDutyFee     int64   `gorm:"not null"`
		MinLoanStampDuty int64   `gorm:"not null"`
		IsActive         bool    `gorm:"not null"`

		Lendings []Lending
	}
)
