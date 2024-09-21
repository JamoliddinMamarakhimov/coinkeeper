package db

import (
	"coinkeeper/configs"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connSrt := fmt.Sprintf(
		`host=%s port=%s user=%s dbname=%s password=%s`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(connSrt), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to database")
	dbConn = db
	return nil
}

func CloseDBConn() error {
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
