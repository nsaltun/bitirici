package db

import (
	"fmt"
	"strings"

	"github.com/nsaltun/bitirici/internal/model"
	"github.com/nsaltun/bitirici/lib/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgDB struct {
	*gorm.DB
}

func NewPostgre() *PgDB {
	v := viper.New()
	v.AutomaticEnv()

	var (
		db  *gorm.DB
		err error
	)

	//check if we are in prod
	//then use the db url from the env
	if strings.ToLower(v.GetString("ENV")) == "prod" {
		//TODO: connect with DB_URL (v.GetString("DB_URL"))

	} else {
		dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			v.GetString("DB_HOST"),
			v.GetString("DB_PORT"),
			v.GetString("DB_USER"),
			v.GetString("DB_PASSWORD"),
			v.GetString("DB_NAME"))
		db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
		if err != nil {
			logging.SugarLog().Fatal("Failed to connect to database", zap.Error(err))
		}
	}

	logging.SugarLog().Info("DB started successfully")
	// Migrate the database
	db.AutoMigrate(&model.User{})
	logging.SugarLog().Info("DB auto migrate is successful")
	return &PgDB{db}
}

func CloseDB(db *gorm.DB) {
	// Close the database connection
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
