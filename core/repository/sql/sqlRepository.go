package sql

import (
	"os"

	"github.com/MiteshSharma/project/core/config"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/jinzhu/gorm"

	// This package is used as mysql driver with gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlRepository struct {
	DB     *gorm.DB
	Log    logger.Logger
	Config config.DatabaseConfig
}

func NewSqlRepository(logger logger.Logger, config config.DatabaseConfig) *SqlRepository {
	sqlRepository := &SqlRepository{
		Log:    logger,
		Config: config,
	}
	sqlRepository.DB = sqlRepository.getDb(config)

	return sqlRepository
}

func (s *SqlRepository) getDb(config config.DatabaseConfig) *gorm.DB {
	var db *gorm.DB
	switch config.Type {
	case "mysql":
		mysqlDb, err := gorm.Open("mysql", config.ConnectionString)
		if err != nil {
			s.Log.Error("Connecting mysql failed due to error ", logger.Error(err))
			os.Exit(1)
		}
		db = mysqlDb
		break
	default:
		break
	}
	return db
}

func (s *SqlRepository) Close() error {
	return s.DB.Close()
}
