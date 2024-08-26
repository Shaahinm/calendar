package config

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		connectToDatabase()
	}
	return db
}

func connectToDatabase() {
	dbType := Envs.DbType

	switch dbType {
	case "sqlite":
		d, err := newSqliteConnection()
		if err != nil {
			panic("failed to connect database")
		}

		db = d
	case "mysql":
		d, err := newMysqlConnection()
		if err != nil {
			panic("failed to connect database")
		}
		db = d
	case "postgres":
		d, err := newPostgresConnection()
		if err != nil {
			panic("failed to connect database")
		}
		db = d
	default:
		panic("Database type not supported")
	}
}

func newPostgresConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Envs.DbHost, Envs.DbUser, Envs.DbPass, Envs.DbName, Envs.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func newMysqlConnection() (*gorm.DB, error) {
	// dsn := "root:@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Envs.DbUser, Envs.DbPass, Envs.DbHost, Envs.DbPort, Envs.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func newSqliteConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(Envs.DbName), &gorm.Config{})
	return db, err
}
