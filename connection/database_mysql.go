package connection

import (
	"fmt"
	"log"
	"time"

	"github.com/aditya3232/url-shortener/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectDatabaseMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", config.CONFIG.DB_USER, config.CONFIG.DB_PASS, config.CONFIG.DB_HOST, config.CONFIG.DB_PORT, config.CONFIG.DB_NAME, config.CONFIG.DB_CHARSET, config.CONFIG.DB_LOC)

	logMode := logger.Silent
	if debug == 1 {
		logMode = logger.Info
		fmt.Println("Database connection string: ", dsn)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	log.Print("Database is connected")
	return db, nil
}

func DatabaseMysql() *gorm.DB {
	return connection.db
}
