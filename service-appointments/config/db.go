package config

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func ConnectToDatabase() {
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
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	return db, err
}

func newMysqlConnection() (*gorm.DB, error) {
	// dsn := "root:@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	builder := strings.Builder{}
	builder.WriteString(Envs.DbUser)
	builder.WriteString(":")
	builder.WriteString("@tcp(")
	builder.WriteString(Envs.DbHost)
	builder.WriteString(":")
	builder.WriteString(Envs.DbPort)
	builder.WriteString(")/")
	builder.WriteString(Envs.DbName)
	builder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	dsn := builder.String()
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func newSqliteConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}
