package connection

import (
	"sync"

	"github.com/aditya3232/url-shortener/config"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

var (
	debug      int = config.CONFIG.DEBUG
	connection Connection
	initOnce   sync.Once
)

func init() {
	initOnce.Do(func() {
		db, err := connectDatabaseMysql()
		if err != nil {
			panic(err)
		}

		connection = Connection{
			db: db,
		}
	})
}

func Close() {
	if connection.db != nil {
		sqlDB, _ := connection.db.DB()
		sqlDB.Close()
		connection.db = nil
	}
}
