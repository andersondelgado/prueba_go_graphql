package mysql

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/andersondelgado/prueba_go_graphql/pkg/config"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDefaultDB() {

	host := config.Environment.DbHost
	//dbDialect := config.Environment.DbConnection
	port := config.Environment.DbPort
	dbUser := config.Environment.DbUser
	dbPassword := config.Environment.DbPassword
	dbName := config.Environment.DbName
	//dbSslMode := config.Environment.DbSslMode

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser,dbPassword,host,port,dbName)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if sqlDB, err := db.DB(); err != nil {
		panic(err)
	} else {
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(10)
	}
	Db = db
}

