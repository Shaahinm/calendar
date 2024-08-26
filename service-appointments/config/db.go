package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func connect() {
	dbType := Envs.DbType

	switch dbType {
	case "sqlite":
		db, err := newSqlliteConnection()
		if err != nil {
			panic("failed to connect database")
		}

		db = db
	case "mysql":
		db, err := newMysqlConnection()
		if err != nil {
			panic("failed to connect database")
		}
		db = db
	case "postgres":
		db, err := newPostgresConnection()
		if err != nil {
			panic("failed to connect database")
		}
		db = db
	default:
		panic("Database type not supported")
	}
}

func newPostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	return db, err
}

func newMysqlConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	return db, err
}

func newSqlliteConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}
