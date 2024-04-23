package config

import (
	"github.com/IsaelVVI/warezapback.git/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initialize logger and mysql
func InitializeMysql() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	// create Database and connect
	dsn := "root:@tcp(127.0.0.1:3306)/warezap"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		logger.Errorf("MYSQL open error: %v", err)
		return nil, err
	}

	// Migrate schemas
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("MYSQL automigration error: %v", err)
		return nil, err
	}

	// Return db
	return db, nil
}
