package main

import (
	"fmt"
	"log"

	"github.com/aburizalpurnama/siesta-test/config"
	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	var config config.Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Failed to load config, unable to decode into struct, %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Database.Host, config.Database.DbUser, config.Database.DbPassword, config.Database.DbName, config.Database.Port, config.Database.SslMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to db, %s", err)
	}

	db.AutoMigrate(
		&model.Account{},
		&model.LendingType{},
		&model.Lending{},
		&model.Repayment{},
	)
	fmt.Println("Success migrate")
}
