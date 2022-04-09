package db

import (
	"api/conf"
	"api/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	var err error
	var dbConfig = conf.Conf.Db

	// dsn := "host=localhost user=postgres password=ji3g4jovul4 dbname=db1 port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Dialects,
		dbConfig.Port,
		"disable")

	// connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode((logger.Info))})
	// Logger: Db.Logger.LogMode(4),
	// })

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(Db.Error)
	}
	Db, _ := Db.DB()

	Db.SetMaxIdleConns(dbConfig.MaxIdle)
	Db.SetMaxOpenConns(dbConfig.MaxOpen)
	logger := utils.Log()
	// Db.SetLogger(logger)
	// Db.LogMode(true)
	logger.Info("postgres connect success")

}
