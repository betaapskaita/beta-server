package databases

import "gorm.io/gorm"

type Database interface {
	Connect() (*gorm.DB, error)
	GetConnection() *gorm.DB
}
