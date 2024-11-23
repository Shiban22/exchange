package models

import (
	"gorm.io/gorm"

	"exchange/utils"
)

type User struct {
	gorm.Model
	ID       int64  `json:"id,string" gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

// 雪花id
func (c *User) BeforeCreate(scope *gorm.DB) error {
	c.ID = utils.GenSnowflakeId()
	return nil
}
