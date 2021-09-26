package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}
