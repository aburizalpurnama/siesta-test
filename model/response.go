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

		Repayments []RepaymentResponse `json:"repayments"`
	}

	RepaymentResponse struct {
		Amount  int64     `json:"amount"`
		DueDate time.Time `json:"dueDate"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)
