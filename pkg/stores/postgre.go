package stores

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreConfig struct {
	Host               string
	DB                 string
	User               string
	Password           string
	Port               int
	MaxRetries         int
	MaxConnections     int
	MaxIdleConnections int
	ConnMaxLifetime    time.Duration
	ConnSleepPeriod    time.Duration
}

type ITableItem interface{}

type TableItem struct {
	gorm.Model
}

type PGStore struct {
	DB *gorm.DB
}

type PGTableStore struct {
	Store *PGStore
}

func NewPostgreDB(dbConfig *PostgreConfig) (db *gorm.DB, err error) {
	connStr := fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s`,
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DB, dbConfig.Password)
	for i := 0; i < dbConfig.MaxRetries; i++ {
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err == nil {
			sqlDB, dbErr := db.DB()
			if dbErr == nil {
				sqlDB.SetMaxOpenConns(dbConfig.MaxConnections)
				sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConnections)
				sqlDB.SetConnMaxLifetime(dbConfig.ConnMaxLifetime * time.Minute)
				return db, nil
			}
		}
		time.Sleep(dbConfig.ConnSleepPeriod)
	}
	return nil, errors.Wrapf(err, "failed to open a connection to postgres database")
}

func NewPostgreStore(dbConfig *PostgreConfig) (*PGStore, error) {
	newStore := &PGStore{}
	db, err := NewPostgreDB(dbConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open a connection to postgres database")
	}

	newStore.DB = db
	return newStore, nil
}

func (s *PGStore) Close() error {
	log.Info().Msg("closing postgre db store")
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
