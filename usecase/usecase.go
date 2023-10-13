package usecase

import (
	"time"

	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/aburizalpurnama/siesta-test/repository"
)

const (
	INITIATED = "INITIATED"
	ON_GOING  = "ON_GOING"
	PAID_OFF  = "PAID_OFF"
	ABORTED   = "ABORTED"
	LATE      = "LATE"
)

type (
	MbtUsecase interface {
		CreateSimulation(reqData model.CreateSimulationRequest) (respData model.CreateSimulationResponse, err error)
	}

	MbtUsecaseImpl struct {
		mbtRepo repository.MbtRepository
	}
)

func NewMbtUsecase(repo repository.MbtRepository) MbtUsecase {
	return &MbtUsecaseImpl{repo}
}

func (u *MbtUsecaseImpl) CreateSimulation(reqData model.CreateSimulationRequest) (respData model.CreateSimulationResponse, err error) {
	// get lending type by id
	lendingType, err := u.mbtRepo.GetLendingTypeById(reqData.LendingTypeId)
	if err != nil {
		return
	}

	monthlyInterest := int64(float64(reqData.Amount) * (lendingType.MonthlyInterest / 100))
	adminFee := int64(float64(reqData.Amount) * (lendingType.AdminFee / 100))
	principalPerMonth := int64(reqData.Amount / int64(reqData.Tenor))

	// create lendings
	lending := model.Lending{
		Date:              reqData.Date,
		Amount:            reqData.Amount,
		Tenor:             reqData.Tenor,
		Principal:         reqData.Amount,
		PrincipalPerMonth: principalPerMonth,
		MonthlyInterest:   monthlyInterest,
		TotalInterest:     monthlyInterest * int64(reqData.Tenor),
		AdminFee:          adminFee,
		LendingPurpose:    reqData.Purpose,
		WithdrawalType:    reqData.WithdrawalType,
		Status:            INITIATED,
		AccountID:         reqData.AccountId,
		LendingTypeID:     reqData.LendingTypeId,
	}

	lendingId, err := u.mbtRepo.CreateLending(lending)
	if err != nil {
		return
	}

	// cek limit yang terpakai untuk dapetin akumulasi pinjaman dan bandingkan dengan ketentuan stampFee
	account, err := u.mbtRepo.GetAccountById(reqData.AccountId)
	if err != nil {
		return
	}

	totalLoan := (account.Limit - account.LimitRemaining) + reqData.Amount
	var stampDutyFee int64
	if totalLoan > lendingType.MinLoanStampDuty {
		stampDutyFee = lendingType.StampDutyFee
	}

	respData.AdminFee = adminFee
	respData.Date = reqData.Date
	respData.LendingType = lendingType.Name
	respData.Day = reqData.Date.Weekday().String()
	respData.MonthlyInterest = monthlyInterest
	respData.Principal = reqData.Amount
	respData.Purpose = reqData.Purpose
	respData.StampDutyFee = stampDutyFee
	respData.WithdrawalType = reqData.WithdrawalType

	var totalBill int64
	repayments := make([]model.RepaymentResponse, reqData.Tenor)
	for i := 0; i < reqData.Tenor; i++ {
		if i != 0 {
			adminFee = 0
		}

		var outstanding int64 = reqData.Amount - (int64(i+1) * principalPerMonth)
		var bill int64 = adminFee + stampDutyFee + monthlyInterest + principalPerMonth
		var dueDate time.Time = reqData.Date.AddDate(0, i+1, 0)
		repayment := model.Repayment{
			DueDate:      dueDate,
			StampDutyFee: stampDutyFee,
			Interest:     monthlyInterest,
			Principal:    principalPerMonth,
			IsPaid:       false,
			Bill:         bill,
			AdminFee:     adminFee,
			Outstanding:  outstanding,
			LendingID:    lendingId,
		}

		_, err = u.mbtRepo.CreateRepayment(repayment)
		if err != nil {
			return
		}

		totalBill += bill
		repayments[i] = model.RepaymentResponse{
			Amount:  bill,
			DueDate: dueDate,
		}
	}

	respData.Total = totalBill

	respData.Repayments = repayments

	return
}
