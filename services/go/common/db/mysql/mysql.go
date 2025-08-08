package mysql

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockery --name=IMySqlDB --case=snake --disable-version-string
type IMySqlDB interface {
	DB() *gorm.DB
	Ping() error
}

type mySQLDBImpl struct {
	db *gorm.DB
}

func NewDB(mySQLConfig *Config) (IMySqlDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLConfig.Username, mySQLConfig.Password, mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(mySQLConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mySQLConfig.MaxOpenConns)

	return &mySQLDBImpl{
		db: db,
	}, nil
}

func (s *mySQLDBImpl) DB() *gorm.DB {
	return s.db
}

func (s *mySQLDBImpl) Ping() error {
	if s.db == nil {
		return errors.New("db is empty")
	}
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	if db == nil {
		return errors.New("db is empty")
	}
	return db.Ping()
}
