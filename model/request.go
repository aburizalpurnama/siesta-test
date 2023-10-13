package model

import "time"

type (
	CreateSimulationRequest struct {
		WithdrawalType string    `json:"withdrawalType"`
		AccountId      uint      `json:"accountId"`
		LendingTypeId  uint      `json:"lendingTypeId"`
		Amount         int64     `json:"amount"`
		Purpose        string    `json:"purpose"`
		Tenor          int       `json:"tenor"`
		Date           time.Time `json:"date"`
	}
)
