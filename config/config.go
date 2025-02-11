package config

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	once   sync.Once
)

func Init() error {
	var err error
	db, err = InitializePostgree()
	if err != nil {
		return fmt.Errorf("error initializing postgree: %v", err)
	}
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	once.Do(func() {
		logger = NewLogger(p)
	})
	return logger
}
