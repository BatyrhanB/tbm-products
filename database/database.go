package database

import (
	"fmt"
	"github.com/BatyrhanB/gin-edu/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func viperEnvVariable(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion for key: %s", key)
	}
	return value
}

var Database *gorm.DB

func Connect() error {
	var err error
	host := viperEnvVariable("POSTGRES_HOST")
	username := viperEnvVariable("POSTGRES_USER")
	password := viperEnvVariable("POSTGRES_PASSWORD")
	databaseName := viperEnvVariable("POSTGRES_DB")
	port := viperEnvVariable("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = Database.AutoMigrate(&model.Product{})
	if err != nil {
		return fmt.Errorf("AutoMigrate failed: %w", err)
	}

	return nil
}
