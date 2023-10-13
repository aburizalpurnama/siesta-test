package usecase

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/aburizalpurnama/siesta-test/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUsecase(t *testing.T) {
	dsn := "host=localhost user=admin password=secret dbname=siesta port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to db, %s", err)
	}

	repo := repository.NewMbtRepo(db)
	usecase := NewMbtUsecase(repo)

	data := model.CreateSimulationRequest{
		WithdrawalType: "CASH",
		AccountId:      1,
		LendingTypeId:  1,
		Amount:         6000000,
		Purpose:        "buy iphone",
		Tenor:          3,
		Date:           time.Now(),
	}

	respData, err := usecase.CreateSimulation(data)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("respData: %v\n", respData)
}
