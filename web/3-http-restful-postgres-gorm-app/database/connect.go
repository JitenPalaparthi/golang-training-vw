package database

import (
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRYS         = 10
	RETRY_DURATION = time.Second * 5
)

func Connect(dsn string) (*gorm.DB, error) {
	var err error
	for i := 1; i <= RETRYS; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		slog.Info("Trying to connecto the database.Try for following number of times-->", i)
		time.Sleep(RETRY_DURATION)
	}
	return nil, err
}
