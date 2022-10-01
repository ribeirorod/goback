package database

import (
	"go-server/cmd/api/config"
	"go-server/cmd/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCon models.DBModel

func InitDB() {
	var cfg = config.GetAppConfig()
	var err error

	DBCon.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Db.Dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("failed to connect database")

	}
}