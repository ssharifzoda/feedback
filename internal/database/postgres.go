package database

import (
	"feedback/pkg/logging"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
}

func NewPostgresGorm() (*gorm.DB, error) {
	Host := viper.GetString("db.host")
	Port := viper.GetUint16("db.port")
	Username := viper.GetString("db.username")
	Password := os.Getenv("DB_PASSWORD")
	DBName := viper.GetString("db.dbname")
	logger := logging.GetLogger()
	connString := fmt.Sprintf("host=%s user=%s password=%v dbname=%s port=%d sslmode=disable TimeZone=Asia/Dushanbe",
		Host, Username, Password, DBName, Port)
	conn, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		logger.Printf("%s GetPostgresConnection -> Open error: ", err.Error())
		return nil, err
	}

	logger.Println("Postgres Connection success: ", Host)

	return conn, nil
}
