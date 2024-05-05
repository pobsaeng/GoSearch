package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:MySQL!234@/test_db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
