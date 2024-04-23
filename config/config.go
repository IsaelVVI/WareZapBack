package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	// Initialize MYSQL
	db, err = InitializeMysql()

	if err != nil {
		return fmt.Errorf("error in init mysql: %v", err)
	}

	// return errors.New("fake Error")
	return nil
}

// Retornando um ponteiro do banco
func GetMYSQL() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}
