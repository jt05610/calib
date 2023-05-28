package calib

import "time"

type Item interface {
	Ref() uint
	CreateTime() *time.Time
	UpdateTime() *time.Time
	DeleteTime() *time.Time
}

type ItemService interface {
	Add(item Item) error
	Remove(item Item) error
	Update(item Item) error
	Get(id uint) (Item, error)
	GetAll() ([]Item, error)
}
