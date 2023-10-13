package repository

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetLendingType(t *testing.T) {
	dsn := "host=localhost user=admin password=secret dbname=siesta port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to db, %s", err)
	}

	repo := NewMbtRepo(db)
	r, err := repo.SelectRepaymentsByAccountId(1)
	fmt.Printf("r: %v\n", r)
	fmt.Printf("err: %v\n", err)

}
