package gorm

import (
	"github.com/jt05610/calib"
	"gorm.io/gorm"
	"time"
)

type point struct {
	gorm.Model
	calID uint
	x     float64
	y     float64
}

func (p *point) Ref() uint {
	return p.ID
}

func (p *point) CreateTime() *time.Time {
	return &p.CreatedAt
}

func (p *point) UpdateTime() *time.Time {
	return &p.UpdatedAt
}

func (p *point) DeleteTime() *time.Time {
	return &p.DeletedAt.Time
}

func (p *point) X() float64 {
	return p.x
}

func (p *point) Y() float64 {
	return p.y
}

func Point(calID uint, x, y float64) calib.Point {
	return &point{calID: calID, x: x, y: y}
}
