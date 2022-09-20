package repo

import (
	"fmt"
	"time"
	"wallet/model"
	"wallet/utils"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	time.Sleep(time.Second * 10)
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func dsn() string {
	host := utils.Env("DB_HOST", "localhost")
	user := utils.Env("DB_USER", "root")
	pass := utils.Env("DB_PASS", "123456")
	dbname := utils.Env("DB_NAME", "arvan")
	port := utils.Env("DB_PORT", "5444")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Customer{}, &model.Wallet{})
}
