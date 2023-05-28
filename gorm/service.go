package gorm

import (
	"github.com/jt05610/calib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) Add(item calib.Calibration) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Remove(item calib.Calibration) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(item calib.Calibration) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(id uint) (calib.Calibration, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetAll() ([]calib.Calibration, error) {
	//TODO implement me
	panic("implement me")
}

func SqliteService(dbName string) calib.ItemService[calib.Calibration] {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Cal{}, &RegResult{}, &Point{})
	if err != nil {
		panic("failed to migrate")
	}
	return &service{db: db}
}
