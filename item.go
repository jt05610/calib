package calib

import "time"

type Item interface {
	Ref() uint
	CreateTime() *time.Time
	UpdateTime() *time.Time
	DeleteTime() *time.Time
}

type ItemService[T Item] interface {
	Add(item T) error
	Remove(item T) error
	Update(item T) error
	Get(id uint) (T, error)
	GetAll() ([]T, error)
}
