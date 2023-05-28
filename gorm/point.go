package gorm

import (
	"gorm.io/gorm"
	"time"
)

type Point struct {
	gorm.Model
	CalID uint
	X     float64
	Y     float64
}

func (p *Point) Ref() uint {
	return p.ID
}

func (p *Point) CreateTime() *time.Time {
	return &p.CreatedAt
}

func (p *Point) UpdateTime() *time.Time {
	return &p.UpdatedAt
}

func (p *Point) DeleteTime() *time.Time {
	return &p.DeletedAt.Time
}
