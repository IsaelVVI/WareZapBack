package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	site_id   *string   `json:"site_id"`
	CreatedAt time.Time `gorm:autoCreateTime json:"created_at"`
	UpdatedAt time.Time `gorm:autoUpdateTime json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
