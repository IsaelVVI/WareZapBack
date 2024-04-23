package schemas

import "gorm.io/gorm"

type Opening struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
	gorm.Model
}
