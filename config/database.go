package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ModelsToMigrate allows modules to register their models for AutoMigrate
var ModelsToMigrate []interface{}

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		AppConfig.DBHost, AppConfig.DBUser, AppConfig.DBPass, AppConfig.DBName, AppConfig.DBPort, AppConfig.DBSSLMode, AppConfig.DBTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             500 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	// AutoMigrate registered models
	if len(ModelsToMigrate) > 0 {
		err = db.AutoMigrate(ModelsToMigrate...)
		if err != nil {
			log.Fatal("❌ AutoMigrate error:", err)
		}
	}

	DB = db
	log.Println("✅ PostgreSQL connected successfully")
}
