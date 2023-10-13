package main

import (
	"fmt"
	"log"

	"github.com/aburizalpurnama/siesta-test/config"
	"github.com/aburizalpurnama/siesta-test/handler"
	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/aburizalpurnama/siesta-test/repository"
	"github.com/aburizalpurnama/siesta-test/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var config config.Config
	getEnv(&config)

	db, err := connectToDB(config)
	if err != nil {
		log.Println(err)
	}

	repo := repository.NewMbtRepo(db)
	usecase := usecase.NewMbtUsecase(repo)
	handler := handler.NewMbtHandler(usecase)

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Post("/accounts/:accountId/lendings", handler.CreateSimulation)
	v1.Post("/lendings/:lendingId/approvals", handler.ApproveLending)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Server.Port)))
}

func getEnv(config *config.Config) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Failed to load config, unable to decode into struct, %v", err)
	}
}

func connectToDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Database.Host, config.Database.DbUser, config.Database.DbPassword, config.Database.DbName, config.Database.Port, config.Database.SslMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to db, %s", err)
	}

	err = db.AutoMigrate(
		&model.Account{},
		&model.LendingType{},
		&model.Lending{},
		&model.Repayment{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate, %s", err)
	}
	fmt.Println("db migration succes")

	err = populateAccount(db)
	if err != nil {
		return nil, err
	}

	err = populateLendingType(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func populateLendingType(db *gorm.DB) error {
	var lenTypes []model.LendingType
	err := db.Find(&lenTypes).Error
	if err != nil {
		return err
	}

	if len(lenTypes) == 0 {
		lenType := model.LendingType{Name: "Murabahah", MonthlyInterest: 1.99, AdminFee: 5, StampDutyFee: 10000, IsActive: true, MinLoanStampDuty: 5000000}
		err = db.Create(&lenType).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func populateAccount(db *gorm.DB) error {
	var accounts []model.Account
	err := db.Find(&accounts).Error
	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		account := model.Account{Name: "Ahmad", Limit: 10000000, LimitRemaining: 10000000}
		err = db.Create(&account).Error
		if err != nil {
			return err
		}
	}

	return nil
}
