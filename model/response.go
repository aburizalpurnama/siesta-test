package model

import "time"

type (
	CreateSimulationResponse struct {
		LendingType     string    `json:"lendingType"`
		WithdrawalType  string    `json:"withdrawalType"`
		Purpose         string    `json:"purpose"`
		Day             string    `json:"day"`
		Date            time.Time `json:"date"`
		Principal       int64     `json:"principal"`
		MonthlyInterest int64     `json:"MonthlyInterest"`
		AdminFee        int64     `json:"adminFee"`
		StampDutyFee    int64     `json:"stampDutyFee"`
		Total           int64     `json:"total"`
		LendingId       uint      `json:"lendingId"`

		Repayments []RepaymentResponse `json:"repayments"`
	}

	RepaymentResponse struct {
		Amount  int64     `json:"amount"`
		DueDate time.Time `json:"dueDate"`
	}

	SelectRepaymentsResponse struct {
		DueDate      time.Time `json:"dueDate"`
		AdminFee     int64     `json:"AdminFee"`
		StampDutyFee int64     `json:"StampDutyFee"`
		Interest     int64     `json:"Interest"`
		Principal    int64     `json:"Principal"`
		Bill         int64     `json:"Bill"`
		Outstanding  int64     `json:"Outstanding"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)
