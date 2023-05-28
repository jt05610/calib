package gorm

import (
	"github.com/jt05610/calib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) Add(item *Cal) error {
	result := s.db.Create(item)

	return result.Error
}

func (s service) Remove(item *Cal) error {
	result := s.db.Delete(item)

	return result.Error
}

func (s service) Update(item *Cal) error {
	result := s.db.Save(item)
	return result.Error
}

func (s service) Get(id uint) (*Cal, error) {
	var ret Cal
	s.db.Model(&Cal{}).Preload("Points").Preload("Result").First(&ret, id)
	if ret.ID == 0 {
		return nil, calib.ErrNotFound
	}
	return &ret, nil
}

func (s service) GetAll() ([]*Cal, error) {
	var ret []*Cal
	s.db.Model(&Cal{}).Preload("Points").Preload("Result").Find(&ret)
	if len(ret) == 0 {
		return nil, calib.ErrNotFound
	}
	return ret, nil
}

func SqliteService(dbName string) calib.ItemService[*Cal] {
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
