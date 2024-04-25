package controllers

import (
	"github.com/IsaelVVI/warezapback.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMYSQL()
}

func GetConfigs() (*gorm.DB, *config.Logger) {
	return db, logger
}
